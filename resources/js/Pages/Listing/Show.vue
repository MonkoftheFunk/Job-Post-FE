<script setup>
import Header from '@/Components/Header.vue'
import {Head} from '@inertiajs/vue3';

defineProps({
    listing: {
        type: Object,
        required: true
    }
})

const dateFormat = (v) => {
    let format = (d) =>
        d.toString().replace(/\w+ (\w+) (\d+) (\d+).*/, "$2 $1, $3");
    return format(new Date(v));
};
</script>

<template>
    <Head :title="listing.title"/>
    <Header/>

    <section class="text-gray-600 body-font overflow-hidden">
        <div class="container px-5 py-24 mx-auto">
            <div class="mb-12">
                <h2 class="text-2xl font-medium text-gray-900 title-font">
                    {{ listing.title }}
                </h2>
                <div class="md:flex-grow mr-8 mt-2 flex items-center justify-start">
                    <span v-for="tag in listing.tags"
                          class="inline-block mr-2 tracking-wide text-indigo-500 text-xs font-medium title-font py-0.5 px-1.5 border border-indigo-500 uppercase">{{ tag.name }}</span>
                </div>
            </div>
            <div class="-my-6">
                <div class="flex flex-wrap md:flex-nowrap">
                    <div v-html="listing.content" class="content w-full md:w-3/4 pr-4 leading-relaxed text-base">
                    </div>
                    <div class="w-full md:w-1/4 pl-4">
                        <img :src="listing.logoUri"
                             :alt="listing.company + ' logo'"
                             class="max-w-full mb-4"
                        >
                        <p class="leading-relaxed text-base">
                            <strong>Location: </strong>{{ listing.location }}<br>
                            <strong>Company: </strong>{{ listing.company }}
                        </p>
                        <a :href="route('listings.apply', {'listing': listing.slug})"
                           class="block text-center my-4 tracking-wide bg-white text-indigo-500 text-sm font-medium title-font py-2 border border-indigo-500 hover:bg-indigo-500 hover:text-white uppercase">Apply
                            Now</a>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>
