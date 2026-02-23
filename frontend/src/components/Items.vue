<script setup>
import { ref, onMounted } from "vue";
import { Trash2, Pencil } from "lucide-vue-next";

const items = ref([]);
const newItem = ref("");
const editingId = ref(null);
const editingTitle = ref("");

const API_URL = "/items";

async function fetchItems() {
    const res = await fetch(API_URL);
    items.value = await res.json();
}

async function addItem() {
    const title = newItem.value.trim();
    if (!title) return;

    await fetch(API_URL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title }),
    });

    newItem.value = "";
    await fetchItems();
}

async function toggleItem(id) {
    await fetch(`${API_URL}/${id}/toggle`, { method: "PATCH" });
    await fetchItems();
}

async function deleteItem(id) {
    if (!confirm("Delete this item?")) return;
    await fetch(`${API_URL}/${id}`, { method: "DELETE" });
    await fetchItems();
}

function startEdit(item) {
    editingId.value = item.ID;
    editingTitle.value = item.Title;
}

async function saveEdit() {
    if (!editingTitle.value.trim()) return;

    await fetch(`${API_URL}/${editingId.value}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title: editingTitle.value }),
    });

    editingId.value = null;
    editingTitle.value = "";
    await fetchItems();
}

function cancelEdit() {
    editingId.value = null;
    editingTitle.value = "";
}

onMounted(fetchItems);
</script>

<template>
    <div class="bg-white rounded-lg shadow-md p-6 mb-6">
        <form @submit.prevent="addItem" class="flex gap-2">
            <input
                v-model="newItem"
                type="text"
                placeholder="Add new item..."
                class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                required
            />
            <button
                type="submit"
                class="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition"
            >
                Add
            </button>
        </form>
    </div>

    <div class="bg-white rounded-lg shadow-md p-6">
        <ul class="space-y-3">
            <li
                v-for="item in items"
                :key="item.ID"
                class="flex items-center justify-between p-3 border-b border-gray-100 last:border-0"
            >
                <div class="flex items-center gap-3">
                    <input
                        type="checkbox"
                        :checked="item.completed"
                        @change="toggleItem(item.ID)"
                        class="w-5 h-5 text-blue-500 rounded focus:ring-blue-500"
                    />
                    <span
                        v-if="editingId !== item.ID"
                        :class="
                            item.completed
                                ? 'line-through text-gray-400'
                                : 'text-gray-700'
                        "
                        class="text-lg"
                    >
                        {{ item.Title }}
                    </span>
                    <input
                        v-else
                        v-model="editingTitle"
                        @keyup.enter="saveEdit"
                        @keyup.escape="cancelEdit"
                        class="flex-1 px-2 py-1 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                </div>
                <div class="flex gap-2">
                    <button
                        v-if="editingId !== item.ID"
                        @click="startEdit(item)"
                        class="px-3 py-1 text-sm text-blue-500 hover:text-blue-700"
                    >
                        <Pencil />
                    </button>
                    <button
                        v-if="editingId === item.ID"
                        @click="saveEdit"
                        class="px-3 py-1 text-sm text-green-500 hover:text-green-700"
                    >
                        Save
                    </button>
                    <button
                        v-if="editingId === item.ID"
                        @click="cancelEdit"
                        class="px-3 py-1 text-sm text-gray-500 hover:text-gray-700"
                    >
                        Cancel
                    </button>
                    <button
                        @click="deleteItem(item.ID)"
                        class="px-3 py-1 text-sm text-red-500 hover:text-red-700"
                    >
                        <Trash2 />
                    </button>
                </div>
            </li>
        </ul>
        <p v-if="items.length === 0" class="text-gray-500 text-center py-8">
            No items yet. Add some!
        </p>
    </div>
</template>
