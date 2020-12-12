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
      <button class="btn edit-btn">
        編集する
      </button>
      <button
        class="btn complete-btn"
        @click="complete(task.id)"
      >
        完了
      </button>
    </div>
    <div
      class="item-wrapper"
    >
      <input v-model="addInputText" type="text" class="add-input-field">
      <button class="btn add-btn" @click="add">
        + タスクを追加する
      </button>
    </div>
  </div>
</template>

<script>
import * as tasksApi from '../api/tasks'

const STATUS_TODO = 0 // 未完了
const STATUS_DONE = 1 // 完了

const DUMMY_LIST = [
  {
    id: 1,
    name: 'タスク1',
    status: STATUS_TODO,
  },
  {
    id: 2,
    name: 'タスク2',
    status: STATUS_TODO,
  },
  {
    id: 3,
    name: 'タスク3',
    status: STATUS_TODO,
  },
]

export default {
  name: 'TodoList',

  data () {
    return {
      list: DUMMY_LIST,
      addInputText: '',
    }
  },

  computed: {
    nextId () {
      return Math.max(...this.list.map(l => l.id)) + 1
    },
    listToShow () {
      return this.list.filter(l => l.status === STATUS_TODO);
    }
  },

  created() {
    this.getAll()
  },

  methods: {
    async getAll() {
      try {
        const res = await tasksApi.getAll()
        this.list = await res.json();
      } catch (err) {
        console.error(err)
      }
    },
    add() {
      if (!this.addInputText) {
        return;
      }

      this.list.push({
        id: this.nextId,
        name: this.addInputText,
        status: STATUS_TODO,
      })

      this.addInputText = ''
    },
    complete(id) {
      // $setしなくても反映されるようになった？
      this.list.find(l => l.id === id).status = STATUS_DONE;
    }
  },
}
</script>

<style scoped>
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
  height: 32px;
  margin-top: 20px;
  padding: 10px;
  border: 1px solid #333;
  font-size: 16px;
}
.item-wrapper > div,
.item-wrapper > button {
  height: 32px;
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
