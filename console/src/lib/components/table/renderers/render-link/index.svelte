<script>
  export let href = ''
  export let text = null
  export let className = ''
  export let external = true

  let props = {}
  let hrefWithPrefix = ''
  let value = ''

  $: {
    props = external ? { target: '_blank', rel: 'noopener noreferrer' } : {}

    if (className) {
      props.class = className
    }

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

    hrefWithPrefix = prefix + href

    value = text || href
  }
</script>

{#if value}
  <a href={hrefWithPrefix} {...props}>{value}</a>
{:else}
  ''
{/if}
