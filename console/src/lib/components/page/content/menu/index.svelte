<script>
  import { pageContentOffsetX } from '../../../../stores'
  import { query } from '../../../../actions'
  import { medium, large } from '../../../../styles/breakpoints'

  let gutterW = 180

  let mediaSize = 'small'
  $: isMedium = query(medium)
  $: isLarge = query(large)

  $: {
    mediaSize = $isLarge ? 'large' : $isMedium ? 'medium' : 'small'

    gutterW = 180

    if (mediaSize !== 'large') {
      gutterW = 22
    }
  }
</script>

<div
  class="page-content"
  style:--width-local="calc(100% - {gutterW * 2 + $pageContentOffsetX}px)"
  style:--padding-local="0 {gutterW}px"
  style:--margin-left-local={$pageContentOffsetX + 'px'}
>
  <slot />
</div>

<style>
  .page-content {
    margin: 48px 0 96px 0;
    display: flex;
    flex-direction: column;
    align-items: center;

    min-height: calc(100% - 194px);

    width: var(--width-local);
    max-width: var(--width-local);
    padding: var(--padding-local);
    margin-left: var(--margin-left-local);

    overflow: hidden;
  }
</style>
