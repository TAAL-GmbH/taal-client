export async function copyTextToClipboard(text) {
  try {
    await navigator.clipboard.writeText(text)
    return { ok: true }
  } catch (error) {
    return { error }
  }
}
