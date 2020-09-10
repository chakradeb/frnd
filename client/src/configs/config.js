const env = process.env.APP_ENV || 'dev';

module.exports = function(){
    const configs = {
        dev: require('./dev.config'),
        prod: require('./prod.config'),
    }

    return configs[env];
}
