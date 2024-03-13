const mix = require('laravel-mix');
const path = require('path');
require('mix-tailwindcss');

mix.alias({
    "@": path.join(__dirname, "resources/js/"),
});

mix.define({
    __VUE_OPTIONS_API__: 'true', // If you are using the options api.
    __VUE_PROD_DEVTOOLS__: 'true', // If you don't want people sneaking around your components in production.
    __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: 'true',
})
mix.js('resources/js/app.js', 'dist')
    .version()
    .sourceMaps()
    .vue()
    .setPublicPath('dist')
    .postCss('resources/css/app.css', 'dist/css', []);

