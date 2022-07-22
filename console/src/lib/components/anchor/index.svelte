<script>
  export let testId = null

  export let href = null
  export let text = null
  export let external = true

  let propObj = {}
  let prefix = ''

  $: if (external) {
    propObj = { target: '_blank', rel: 'noopener noreferrer' }
    try {
      const url = new URL(href)
      if (!url.protocol) {
        prefix = 'https://'
      }
    } catch (e) {
      prefix = 'https://'
    }
  }

  let optProps = {}
  $: {
    if (testId) {
      optProps['data-test-id'] = testId
    }
  }
</script>

<a href={prefix + href} {...propObj} {...optProps}>{text || href}</a>

<style>
  a {
    color: #0094ff;
  }
</style>
