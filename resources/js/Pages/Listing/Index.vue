<script setup>
import {ref} from 'vue';
import {Head} from '@inertiajs/vue3';
import Hero from '@/Components/Hero.vue';
import Header from '@/Components/Header.vue';
import Footer from "@/Components/Footer.vue";

defineProps({
    search: {
        type: String,
    },
    active_tag: {
        type: String,
    },
    tags: {
        type: Array,
        required: true,
    },
    listings: {
        type: Array,
        required: true,
    },
})

const activeClass = ref('bg-indigo-500 text-white')
const normalClass = ref('bg-white text-indigo-500')
</script>

<template>
    <Head title="Listings"/>
    <Header/>

    <Hero :search="search"/>
    <section class="container px-5 py-12 mx-auto">
        <div class="mb-12">
            <div class="flex-justify-center">
                <a v-for="tag in tags" :href="route('listings.index', {'tag': tag.slug})"
                   class="inline-block ml-2 tracking-wide text-xs font-medium title-font py-0.5 px-1.5 border border-indigo-500 uppercase"
                   :class="[tag.slug === active_tag ? activeClass : normalClass]"
                >{{ tag.name }}</a>
            </div>
        </div>
        <div class="mb-12">
            <h2 class="text-2xl font-medium text-gray-900 title-font px-4">All jobs ({{ listings.length }})</h2>
        </div>
        <div class="-my-6">
            <a v-for="listing in listings"
               :href="route('listings.show', {'slug': listing.slug})"
               class="py-6 px-4 flex flex-wrap md:flex-nowrap border-b border-gray-100"
               :class="[listing.is_highlighted ? 'bg-yellow-100 hover:bg-yellow-200' : 'bg-white hover:bg-gray-100']"
            >
                <div class="md:w-16 md:mb-0 mb-6 mr-4 flex-shrink-0 flex flex-col">
                    <img :src="listing.logoUri" :alt="listing.company + ' logo'" class="w-16 h-16 rounded-full object-cover">
                </div>
                <div class="md:w-1/2 mr-8 flex flex-col items-start justify-center">
                    <h2 class="text-xl font-bold text-gray-900 title-font mb-1">{{ listing.title }}</h2>
                    <p class="leading-relaxed text-gray-900">
                        {{ listing.company }} &mdash; <span class="text-gray-600">{{ listing.location }}</span>
                    </p>
                    <p class="leading-relaxed text-gray-900">
                        Applications {{ listing.clicksCount }}
                    </p>
                </div>
                <div class="md:flex-grow mr-8 flex items-center justify-start">
                    <span v-for="tag in listing.tags"
                          class="inline-block ml-2 tracking-wide text-xs font-medium title-font py-0.5 px-1.5 border border-indigo-500 uppercase "
                          :class="[tag.slug === active_tag ? activeClass : normalClass]">
                           {{ tag.name }}
                    </span>
                </div>
                <span class="md:flex-grow flex items-center justify-end">
                    <span>{{ listing.sinceCreated }}</span>
                </span>
            </a>
        </div>
    </section>

    <Footer/>
</template>
