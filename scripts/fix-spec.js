// Corrects known issues in the upstream Statuspage OpenAPI spec that break Go code generation.
const fs = require('fs');

const specPath = 'developer_statuspage_io.json';
const spec = JSON.parse(fs.readFileSync(specPath, 'utf8'));

let fixedTags = 0;
let fixedDateParamTypes = 0;
let fixedObjectTypes = 0;
let fixedDateTimeDefaults = 0;

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

    for (const value of Object.values(node)) walk(value);
}

walk(spec);

fs.writeFileSync(specPath, JSON.stringify(spec, null, 2));
console.log(
    `Deduplicated tags on ${fixedTags} operation(s), ` +
    `fixed ${fixedDateParamTypes} invalid date param type(s), ` +
    `fixed ${fixedObjectTypes} invalid Object type(s), ` +
    `removed ${fixedDateTimeDefaults} invalid date-time default(s).`
);

