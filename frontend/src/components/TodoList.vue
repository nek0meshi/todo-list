<template>
  <div class="list">
    <div
      v-for="task in listToShow"
      :key="task.id"
      class="item-wrapper"
    >
      <div class="item-text">
        {{ task.name }}
      </div>
      <button
        class="btn edit-btn button is-light"
        @click="edit(task.id)"
      >
        編集する
      </button>
      <button
        class="btn complete-btn button is-primary"
        @click="complete(task.id)"
      >
        完了
      </button>
    </div>
    <div
      class="item-wrapper"
    >
      <input v-model="addInputText" type="text" class="add-input-field">
      <button class="btn add-btn button is-success" @click="add">
        + タスクを追加する
      </button>
    </div>
    <EditModal
      :show="showEditModal"
      :text="editText"
      @save="saveEdit"
      @close="showEditModal = false"
    />
  </div>
</template>

<script>
import * as tasksApi from '../api/tasks'
import '../../node_modules/bulma/css/bulma.css'
import EditModal from './EditModal.vue'

const STATUS_TODO = 0 // 未完了
const STATUS_DONE = 1 // 完了

export default {
  name: 'TodoList',

  components: {
    EditModal,
  },

  data () {
    return {
      list: [],
      addInputText: '',
      showEditModal: false,
      editId: null,
    }
  },

  computed: {
    nextId () {
      return Math.max(...this.list.map(l => l.id)) + 1
    },
    listToShow () {
      return this.list.filter(l => l.status === STATUS_TODO);
    },
    editText() {
      return this.list.find(({ id }) => id === this.editId)?.name
    }
  },

  created() {
    this.getAll()
  },

  methods: {
    edit(id) {
      this.editId = id
      this.showEditModal = true
    },
    saveEdit({ text }) {
      this.update(this.editId, text)
      this.showEditModal = false
    },
    async getAll() {
      try {
        const res = await tasksApi.getAll()
        this.list = await res.json();
      } catch (err) {
        console.error(err)
      }
    },
    async add() {
      if (!this.addInputText) {
        return
      }

      try {
        const res = await tasksApi.store({
          name: this.addInputText,
        })
        this.addInputText = ''
        this.getAll()
      } catch (err) {
        console.error(err)
      }
    },
    async complete(id) {
      try {
        const res = await tasksApi.complete(id)
        this.getAll()
      } catch (err) {
        console.error(err)
      }
    },
    async update(id, text) {
      try {
        const res = await tasksApi.update(id, { name: text })
        this.getAll()
      } catch (err) {
        console.error(err)
      }
    }
  },
}
</script>

<style scoped lang="scss">
.list {
  display: flex;
  flex-direction: column;
  margin: 0 auto;
  width: 540px;
}

.item-wrapper {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  height: 54px;
  margin-top: 20px;
  padding: 10px;
  border: 1px solid #333;
  font-size: 16px;
  > div,
  > button {
    height: 32px;
  }
}
.item-text {
  flex: 1;
  font-size: 20px;
  text-align: left;
}
.btn {
  margin-left: 10px;
}
.add-input-field {
  text-align: left;
  font-size: 16px;
}
</style>
