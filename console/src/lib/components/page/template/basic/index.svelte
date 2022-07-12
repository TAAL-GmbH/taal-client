<script>
  import { headerHeight } from '../../../../stores'
  import Header from '../../../header/index.svelte'
  import Footer from '../../../footer/index.svelte'
  import ContentBasic from '../../content/basic/index.svelte'
</script>

<Header showLinks={false} showActions={false} />

<div class="content-container" style:--top-local={$headerHeight + 'px'}>
  <ContentBasic>
    <slot />
  </ContentBasic>

  <Footer />
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
