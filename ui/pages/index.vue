<template>
  <div>
    <section class="hero is-primary">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">
            {{ title }}
          </h1>
          <h2 class="subtitle">
            {{ subtitle }}
          </h2>
        </div>
      </div>
    </section>
    <div class="container">
      <a class="button" @click="showModal">Show</a>
    </div>
    <div class="container">
      <nuxt-link :to="note">
        Notepad
      </nuxt-link>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: 'Cool dude!',
      title: 'Title',
      subtitle: 'Subtitle',
      second: '/second',
      note: '/note',
      about: '/about',
      login: '/login'
    }
  },
  methods: {
    showModal() {
      alert('Cool!')
    }
  }
}
</script>
