<script>
  import { headerHeight } from '../../../../stores'
  import Header from '../../../header/index.svelte'
  import Footer from '../../../footer/index.svelte'
  import ContentBasic from '../../content/basic/index.svelte'

  export let testId = null

  let optProps = {}
  $: {
    if (testId) {
      optProps['data-test-id'] = testId
    }
  }
</script>

<Header showLinks={false} showActions={false} testId="header" />

<div
  class="content-container"
  style:--top-local={$headerHeight + 'px'}
  {...optProps}
>
  <ContentBasic>
    <slot />
  </ContentBasic>

  <Footer testId="footer" />
</div>

<style>
  .content-container {
    position: absolute;
    top: var(--top-local);
    width: 100%;
    height: calc(100% - var(--top-local));
    overflow-x: hidden;
    overflow-y: auto;
  }
</style>
