<script>
  import { mediaSize, pageContentOffsetX } from '../../../../stores'

  $: gutterW = $mediaSize !== 'large' ? 22 : 180
  $: marginTop = $mediaSize === 'small' ? 32 : 48
</script>

<div
  class="page-content"
  style:--width-local="calc(100% - {gutterW * 2 + $pageContentOffsetX}px)"
  style:--padding-local="0 {gutterW}px"
  style:--margin-top-local={marginTop + 'px'}
  style:--margin-left-local={$pageContentOffsetX + 'px'}
  style:--min-height-local="calc(100% - {marginTop + 96 + 50}px)"
>
  <slot />
</div>

<style>
  .page-content {
    margin: var(--margin-top-local) 0 96px 0;
    display: flex;
    flex-direction: column;
    align-items: center;

    min-height: var(--min-height-local);

    width: var(--width-local);
    max-width: var(--width-local);
    padding: var(--padding-local);
    margin-left: var(--margin-left-local);

    overflow: hidden;
  }
</style>
