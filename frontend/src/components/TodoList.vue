<template>
  <div class="list">
    <div
      v-for="task in list"
      :key="task.id"
      class="item-wrapper"
    >
      <div class="item-text">
        {{ task.name }}
      </div>
      <button class="btn edit-btn">
        編集する
      </button>
      <button class="btn complete-btn">
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
const DUMMY_LIST = [
  {
    id: 1,
    name: 'タスク1',
    status: 0, // 0: 未完了, 1: 完了
  },
  {
    id: 2,
    name: 'タスク2',
    status: 0, // 0: 未完了, 1: 完了
  },
  {
    id: 3,
    name: 'タスク3',
    status: 0, // 0: 未完了, 1: 完了
  },
];

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
      return Math.max(this.list.map(l => l.id))
    },
  },

  methods: {
    add() {
      if (!this.addInputText) {
        return;
      }
      this.list.push({
        id: this.nextId,
        name: this.addInputText,
        status: 0,
      })

      this.addInputText = ''
    },
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
