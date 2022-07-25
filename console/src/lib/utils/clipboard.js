import { success, failure } from './notifications'
import i18n from '../i18n'

let t = () => {}
i18n.subscribe((value) => {
  t = value.t
})

export async function copyTextToClipboard(text) {
  try {
    await navigator.clipboard.writeText(text)
    success(t('notifications.copy-success'))
    return { ok: true }
  } catch (error) {
    failure(t('notifications.copy-failure'))
    return { error }
  }
}
