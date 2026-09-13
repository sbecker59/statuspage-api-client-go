// Corrects known issues in the upstream Statuspage OpenAPI spec that break Go code generation.
const fs = require('fs');

const specPath = 'developer_statuspage_io.json';
const spec = JSON.parse(fs.readFileSync(specPath, 'utf8'));

let fixedTags = 0;
let fixedDateParamTypes = 0;
let fixedObjectTypes = 0;
let fixedDateTimeDefaults = 0;
let fixedComponentMapTypes = 0;
let fixedArrayTypes = 0;
let fixedPositionTypes = 0;
let fixedComponentGroupDescriptions = 0;
let fixedPageImageFieldTypes = 0;
let fixedMetricsProviderPageIdTypes = 0;

// Operations tagged with more than one tag get duplicated across multiple
// generated Go files (one per tag), causing "redeclared in this block" errors.
for (const methods of Object.values(spec.paths || {})) {
    for (const op of Object.values(methods)) {
        if (!op || typeof op !== 'object') continue;
        if (Array.isArray(op.tags) && op.tags.length > 1) {
            op.tags = [op.tags[0]];
            fixedTags++;
        }
    }
}

// Recursively walk every schema object in the spec to fix invalid/undefined
// type names and defaults that don't survive Go code generation.
function walk(node) {
    if (Array.isArray(node)) {
        for (const item of node) walk(item);
        return;
    }
    if (!node || typeof node !== 'object') return;

    // "type": "PartialStartDate"/"PartialEndDate" are not valid OpenAPI types and
    // have no matching schema definition, so the generator emits an undefined Go type.
    if (node.type === 'PartialStartDate' || node.type === 'PartialEndDate') {
        node.type = 'string';
        node.format = 'date';
        fixedDateParamTypes++;
    }

    // "type": "Object" (capitalized) is not a valid OpenAPI type, so the generator
    // emits it verbatim as an undefined Go type "Object" instead of a free-form object.
    if (node.type === 'Object') {
        node.type = 'object';
        fixedObjectTypes++;
    }

    // Defaults on date-time fields are emitted as raw string literals assigned to a
    // time.Time variable, which does not compile. Drop the default in that case.
    if (node.format === 'date-time' && typeof node.default === 'string') {
        delete node.default;
        fixedDateTimeDefaults++;
    }

    // Component map objects in incident request bodies are documented with dummy
    // property names (e.g. "7bsz4wf6bh17", "xw9j0mrxmgrv"), causing the generator
    // to emit rigid structs instead of map[string]string. Convert them to additionalProperties.
    if (
        node.type === 'object' &&
        node.properties &&
        typeof node.description === 'string' &&
        node.description.includes('Map of status changes to apply to affected components')
    ) {
        delete node.properties;
        node.additionalProperties = { type: 'string' };
        fixedComponentMapTypes++;
    }

    // Properties typed as string but having an array example should be arrays of strings.
    if (node.type === 'object' && node.properties) {
        for (const [propName, prop] of Object.entries(node.properties)) {
            if (prop && prop.type === 'string' && Array.isArray(prop.example)) {
                prop.type = 'array';
                prop.items = { type: 'string' };
                fixedArrayTypes++;
            }
        }
    }

    // Fix position type in GroupComponent from string to integer (int32)
    if (
        node.type === 'object' &&
        node.properties &&
        node.description === 'Get a component group' &&
        node.properties.position &&
        node.properties.position.type === 'string'
    ) {
        node.properties.position = {
            type: 'integer',
            format: 'int32',
            description: 'Order the component group will appear on the page',
        };
        fixedPositionTypes++;
    }

    for (const value of Object.values(node)) walk(value);
}

walk(spec);

// The spec documents "description" as a top-level sibling of "component_group" in the
// create/update component group request bodies, but the live API only persists it when
// nested inside "component_group". Move it there so the generated client sends it correctly.
for (const schemaName of ['postPagesPageIdComponentGroups', 'patchPagesPageIdComponentGroups', 'putPagesPageIdComponentGroups']) {
    const schema = spec.components?.schemas?.[schemaName];
    const nestedSchema = schema?.properties?.component_group;
    if (schema?.properties?.description && nestedSchema?.properties && !nestedSchema.properties.description) {
        nestedSchema.properties.description = schema.properties.description;
        delete schema.properties.description;
        fixedComponentGroupDescriptions++;
    }
}

// The live API returns either a string URL or an attachment object for these Page image
// fields, but the spec documents them as plain strings, breaking Go JSON unmarshaling.
// Make them free-form so the generator emits interface{} instead of string.
const pageProps = spec.components?.schemas?.Page?.properties;
if (pageProps) {
    for (const field of ['favicon_logo', 'transactional_logo', 'hero_cover', 'email_logo', 'twitter_logo']) {
        if (pageProps[field]?.type === 'string') {
            pageProps[field] = {};
            fixedPageImageFieldTypes++;
        }
    }
}

// The live API returns the usual alphanumeric page ID string for MetricsProvider.page_id,
// but the spec documents it as an int32, breaking Go JSON unmarshaling.
const metricsProviderProps = spec.components?.schemas?.MetricsProvider?.properties;
if (metricsProviderProps?.page_id?.type === 'integer') {
    metricsProviderProps.page_id = { type: 'string' };
    fixedMetricsProviderPageIdTypes++;
}

fs.writeFileSync(specPath, JSON.stringify(spec, null, 2));
console.log(
    `Deduplicated tags on ${fixedTags} operation(s), ` +
    `fixed ${fixedDateParamTypes} invalid date param type(s), ` +
    `fixed ${fixedObjectTypes} invalid Object type(s), ` +
    `removed ${fixedDateTimeDefaults} invalid date-time default(s), ` +
    `fixed ${fixedComponentMapTypes} component map type(s), ` +
    `fixed ${fixedArrayTypes} invalid string-array type(s), ` +
    `fixed ${fixedPositionTypes} invalid position type(s), ` +
    `moved ${fixedComponentGroupDescriptions} component group description field(s), ` +
    `loosened ${fixedPageImageFieldTypes} Page image field type(s), ` +
    `fixed ${fixedMetricsProviderPageIdTypes} MetricsProvider page_id type(s).`
);

