<template>
  <div class="edit-modal modal" :class="{ 'is-active': show }">
    <div class="modal-background"></div>
    <div class="modal-card">
      <header class="modal-card-head">
        <p class="modal-card-title">タスクを編集</p>
        <button class="delete" aria-label="close" @click="close"></button>
      </header>
      <section class="modal-card-body">
        <input type="text" class="input" v-model="innerText" />
      </section>
      <footer class="modal-card-foot">
        <button class="button is-success" @click="save">保存</button>
        <button class="button" @click="close">キャンセル</button>
      </footer>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    show: {
      type: Boolean,
      default: false,
    },
    text: {
      type: String,
      default: '',
    },
  },

  data() {
    return {
      innerText: '',
    }
  },

  watch: {
    show(val) {
      if (val) {
        this.innerText = this.text
      }
    },
  },

  methods: {
    save() {
      this.$emit('save', {
        text: this.innerText,
      })
    },
    close() {
      this.$emit('close')
    },
  },
}
</script>
