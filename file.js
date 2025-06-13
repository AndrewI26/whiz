const fs = require("node:fs/promises")
const path = require('path');

async function read_file() {
    try {
        const log_path = path.join(__dirname, "config", "log_config json")
        const stream = await (await fs.open(log_path)).readFile();
        const config = JSON.parse(parse);

        console.log("[File contents] -> \n %s", config.log_prefix);
    }
    catch (err) {
        console.error("Error occured while reading file: %o", err);
    }
}

module.exports = read_file;
