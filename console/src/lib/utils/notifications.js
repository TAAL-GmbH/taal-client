import { toast } from '@zerodevx/svelte-toast'
import Notofication from '../components/notification/index.svelte'

export const success = (m, opts = {}) => {
  const { theme = {}, rest } = opts
  toast.push({
    component: {
      src: Notofication,
      props: { status: 'success', title: m },
    },
    theme: {
      '--toastBarBackground': '#6EC492',
      ...theme,
    },
    ...rest,
  })
}

export const failure = (m, opts = {}) => {
  const { theme = {}, rest } = opts
  toast.push({
    component: {
      src: Notofication,
      props: { status: 'failure', title: m },
    },
    theme: {
      '--toastBarBackground': '#FF344C',
      ...theme,
    },
    ...rest,
  })
}
