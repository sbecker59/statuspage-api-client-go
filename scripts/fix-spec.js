// Corrects known issues in the upstream Statuspage OpenAPI spec that break Go code generation.
const fs = require('fs');

const specPath = 'developer_statuspage_io.json';
const spec = JSON.parse(fs.readFileSync(specPath, 'utf8'));

let fixedTags = 0;
let fixedDateTypes = 0;

for (const methods of Object.values(spec.paths || {})) {
    for (const op of Object.values(methods)) {
        if (!op || typeof op !== 'object') continue;

        // Operations tagged with more than one tag get duplicated across multiple
        // generated Go files (one per tag), causing "redeclared in this block" errors.
        if (Array.isArray(op.tags) && op.tags.length > 1) {
            op.tags = [op.tags[0]];
            fixedTags++;
        }

        for (const param of op.parameters || []) {
            const schema = param && param.schema;
            if (schema && (schema.type === 'PartialStartDate' || schema.type === 'PartialEndDate')) {
                schema.type = 'string';
                schema.format = 'date';
                fixedDateTypes++;
            }
        }
    }
}

fs.writeFileSync(specPath, JSON.stringify(spec, null, 2));
console.log(`Deduplicated tags on ${fixedTags} operation(s), fixed ${fixedDateTypes} invalid date param type(s).`);
