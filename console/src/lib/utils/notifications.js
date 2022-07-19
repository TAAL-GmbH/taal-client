import { toast } from '@zerodevx/svelte-toast'

export const success = (m, opts = {}) =>
  toast.push(m, {
    theme: {
      '--toastBackground': '#6EC492',
      '--toastColor': 'white',
      '--toastBarBackground': 'green',
    },
    ...opts,
  })

export const failure = (m, opts = {}) =>
  toast.push(m, {
    theme: {
      '--toastBackground': '##FF344C',
      '--toastColor': 'white',
      '--toastBarBackground': 'tomato',
    },
    ...opts,
  })
