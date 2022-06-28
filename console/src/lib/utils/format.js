export const link = (href, text = null, className, external = true) => {
  const propObj = external
    ? { target: '_blank', rel: 'noopener noreferrer' }
    : {}
  let prefix = ''
  if (external) {
    try {
      const url = new URL(href)
      if (!url.protocol) {
        prefix = 'https://'
      }
    } catch (e) {
      prefix = 'https://'
    }
  }
  return `<a href='${prefix + href}'${
    external ? " target='_blank' rel='noopener noreferrer'" : ''
  }${className ? ` class='${className}'` : ''}>${text || href}</a>`
}
