<script>
  import Header from '../../../header/index.svelte'
  import Footer from '../../../footer/index.svelte'
  import ContentBasic from '../../content/basic/index.svelte'
</script>

<Header showLinks={false} showActions={false} />

<div class="content-container">
  <ContentBasic>
    <slot />
  </ContentBasic>

  <Footer />
</div>

<style>
  .content-container {
    position: absolute;
    top: 80px;
    width: 100%;
    height: calc(100% - 80px);
    overflow-x: hidden;
    overflow-y: auto;
  }
</style>
