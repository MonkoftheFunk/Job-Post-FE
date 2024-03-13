import { createApp, h } from "vue";
import { createInertiaApp, Link } from "@inertiajs/vue3";
import { ZiggyVue } from 'ziggy-js';

createInertiaApp({
    resolve: (name) => require(`./Pages/${name}`),
    setup({ el, App, props, plugin}) {
        createApp({ render: () => h(App, props) })
            .use(plugin)
            .use(ZiggyVue, Ziggy)
            .mount(el);
    }
})

/**
 * Added to support route(... function in Vue files from laravel,
 * also useful to abstract the routes
 * @type {{routes: {"listings.index": {methods: string[], uri: string}, "listings.show": {methods: string[], bindings: {listing: string}, uri: string, parameters: string[]}, "listings.apply": {methods: string[], bindings: {listing: string}, uri: string, parameters: string[]}}, defaults: {}}}
 */
const Ziggy = {
    "url": "",
    "port": 9180,
    "defaults": {},
    "routes": {
        "listings.index": {
            "uri": "\/",
            "methods": ["GET", "HEAD"]
        },
        "listings.show": {
            "uri": "l\/{listing}",
            "methods": ["GET", "HEAD"],
            "parameters": ["listing"],
            "bindings": {
                "listing": "slug"
            }
        },
        "listings.apply": {
            "uri": "l\/{listing}\/apply",
            "methods": ["GET", "HEAD"],
            "parameters": ["listing"],
            "bindings": {
                "listing": "slug"
            }
        },
        "login": {
            "uri": "login",
            "methods": ["GET", "HEAD"]
        },
        "dashboard": {
            "uri": "dashboard",
            "methods": ["GET", "HEAD"]
        },
    }
};

export { Ziggy };
