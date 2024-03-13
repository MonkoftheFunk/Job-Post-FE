//todo
/*const props = defineProps({
stripe_key: {
type: String,
required: true
}
})

const addScript = (s) => {
let script = document.createElement('script')
script.setAttribute('src', s)
document.head.appendChild(script)
}
addScript('https://js.stripe.com/v3/')

const token = ref(null)
const stripe = ref(null)
const elements = ref(null)

onMounted(() => {
stripe.value = Stripe(props.stripe_key);
})*/

//onMounted(() => {
//    const stripe = Stripe("pk_test_A7jK4iCYHL045qgjjfzAfPxu");

/*const elements = stripe.elements();
const cardElement = elements.create('card', {
classes: {
base: 'StripeElement rounded-md shadow-sm bg-white px-2 py-3 border border-gray-300 focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 block mt-1 w-full'
}
});
cardElement.mount('#card-element');*/
//})
/*c*/

/*
document.getElementById('form_submit').addEventListener('click', async (e) => {
// prevent the submission of the form immediately
e.preventDefault();

const { paymentMethod, error } = await stripe.createPaymentMethod(
'card', cardElement, {}
);

if (error) {
alert(error.message);
} else {
// card is ok, create payment method id and submit form
document.getElementById('payment_method_id').value = paymentMethod.id;
document.getElementById('payment_form').submit();
}
})*/
