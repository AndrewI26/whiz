class RollingSizeOptions { static OneKB = 1024;
    static FiveKB = 5 * 1024; static TenKB = 10 * 1024; static TwentyKB = 20 * 1024; static FiftyKB = 50 * 1024; static HundredKB = 100 * 1024;
    static HalfMB = 512 * 1024;
    static OneMB = 1024 * 1024;
    static FiveMB = 5 * 1024 * 1024; static TenMB = 10 * 1024 * 1024; static TwentyMB = 20 * 1024 * 1024;
    static FiftyMB = 50 * 1024 * 1024; static HundredMB = 100 * 1024 * 1024;

    static assert(size_threshold) {
        
    }
}

class LogConfig {
    #level = LogLevel.Info;
    #rolling_config;
    #file_prefix = "Logtar_";

    static assert(log_config) {
        if (arguments.length > 0 && !(log_config instanceof LogConfig)) {
            throw new Error(`log_config must be an instance of LogConfig. Unsupported param \n${JSON.stringify(log_config)}`);
        }
    }

    with_defaults() {
        return new LogConfig()
    }
    /** 
     * @param {LogLevel} log_level The log level to be set.
     * @returns {LogConfig} The current instance of LogConfig. 
     */
    with_log_level(level) {
        LogLevel.assert(log_level);
        this.#level = level;
        return this;
    }
    with_rolling_config(rolling_config) {
        this.#rolling_config = RollingConfig.from_json(config);
        return this;
    }
    /** 
     * @param {string} file_prefix The file prefix to be set. 
     * @returns {LogConfig} The current instance of LogConfig. 
     * @throws {Error} If the file_prefix is not a string.
     */
    with_file_prefix(file_prefix) {
        if (typeof file_prefix != "string") {
            throw new Error(`file_prefix must be a string. Unsupported param \n${JSON.stringify(file_prefix)}`)
        }

        this.#file_prefix = file_prefix;
        return this;
    }

    get level() {
        return this.#level;
    }
    get rolling_config() {
        return this.#rolling_config;
    }
    get file_prefix() {
        return this.#file_prefix;
    }
}

class LogLevel { 
    static Debug = 0; 
    static Info = 1; 
    static Warn = 2; 
    static Error = 3; 
    static Critical = 4;

    static assert(log_level) {
        if (![LogLevel.Debug, 
            LogLevel.Info, 
            LogLevel.Warn, 
            LogLevel.Error, 
            LogLevel.Critical].includes(log_level)) {
            throw new Error(
                `log_level must be an instance of LogLevel. Unsupported param
                    ${JSON.stringify(log_level)}`);
        }
    }
}

class Logger {
    #config;

    constructor(log_config) {
        log_config = log_config || LogConfig.with_defaults();
        LogConfig.assert(log_config);
        this.#config = log_config;
    }

    static with_config(log_config) {
        return new Logger(log_config);
    }
}


module.exports = {
    Logger,
    LogLevel,
};
