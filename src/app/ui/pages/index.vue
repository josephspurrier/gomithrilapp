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
    <br />
    <div class="container">
      <n-link :to="note" data-cy="notepad-link">
        Click here to access your Notepad.
      </n-link>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: 'Welcome',
      subtitle: 'Login was successful',
      note: '/note'
    }
  },
  head() {
    return {
      title: this.title
    }
  },
  middleware: 'onlyAuthenticated'
}
</script>
