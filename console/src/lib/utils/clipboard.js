import { success, failure } from './notifications'

export async function copyTextToClipboard(text) {
  try {
    await navigator.clipboard.writeText(text)
    success('Successfully copied.', { duration: 2000 })
    return { ok: true }
  } catch (error) {
    failure('Failed to copy.', { duration: 2000 })
    return { error }
  }
}
