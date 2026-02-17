<script setup lang="ts">
import { format, isToday } from 'date-fns'
import type {ConversationResponse} from "$shared/types"
// import type { Mail } from '~/types'

const props = defineProps<{
  conversation: ConversationResponse[]
}>()
const selectedConversation = defineModel<ConversationResponse | null>()
const mailsRefs = ref<Record<number, Element | null>>({})

// const selectedMail = defineModel<Mail | null>()

// watch(selectedMail, () => {
//   if (!selectedMail.value) {
//     return
//   }
//   const ref = mailsRefs.value[selectedMail.value.id]
//   if (ref) {
//     ref.scrollIntoView({ block: 'nearest' })
//   }
// })

// defineShortcuts({
//   arrowdown: () => {
//     const index = props.mails.findIndex((mail: Mail) => mail.id === selectedMail.value?.id)

//     if (index === -1) {
//       selectedMail.value = props.mails[0]
//     } else if (index < props.mails.length - 1) {
//       selectedMail.value = props.mails[index + 1]
//     }
//   },
//   arrowup: () => {
//     const index = props.mails.findIndex((mail: Mail) => mail.id === selectedMail.value?.id)

//     if (index === -1) {
//       selectedMail.value = props.mails[props.mails.length - 1]
//     } else if (index > 0) {
//       selectedMail.value = props.mails[index - 1]
//     }
//   }
// })
</script>

<template>
  <div class="overflow-y-auto divide-y divide-default">
    <div
      v-for="(mail, index) in conversation"
      :key="index"
      :ref="(el) => { mailsRefs[mail.id] = el as Element | null }"
    >
      <div
        class="p-4 sm:px-6 text-sm cursor-pointer border-l-2 transition-colors"
          @click="selectedConversation = mail"
      >
        <div class="flex items-center justify-between" >
          <div class="flex items-center gap-3">
            {{ mail.organization?.name }}
          </div>

          <span>{{ isToday(new Date(mail.updatedAt)) ? format(new Date(mail.updatedAt), 'HH:mm') : format(new Date(mail.updatedAt), 'dd MMM') }}</span>
        </div>
        <p class="truncate" >
          <!-- {{ mail.subject }} -->
        </p>
        <p class="text-dimmed line-clamp-1">
          <!-- {{ mail.body }} -->
        </p>
      </div>
    </div>
  </div>
</template>
