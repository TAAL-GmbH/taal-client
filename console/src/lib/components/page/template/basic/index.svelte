<script>
  import { headerHeight } from '../../../../stores'
  import Header from '../../../header/index.svelte'
  import Footer from '../../../footer/index.svelte'
  import ContentBasic from '../../content/basic/index.svelte'

  export let testId = null
</script>

<Header showLinks={false} showActions={false} testId="header" />

<div
  class="content-container"
  data-test-id={testId}
  style:--top-local={$headerHeight + 'px'}
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
