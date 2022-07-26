<script>
  import { createEventDispatcher } from 'svelte'
  import Button from '../button/index.svelte'
  import { getBtnData } from './utils'

  const dispatch = createEventDispatcher()

  export let testId = null

  let type = 'page'

  export let name = ''
  export let value = 1
  export let pageSize = 10
  export let totalItems = 0
  export let siblingCount = 1
  export let boundaryCount = 1
  export let hasBoundaryRight = true
  export let dataSize = 10

  let totalPages = 1
  let isLastPage = false
  let btnData = []

  $: {
    totalPages = Math.max(1, Math.ceil(totalItems / pageSize))
    isLastPage = dataSize < pageSize

    btnData = getBtnData(
      totalPages,
      value,
      isLastPage,
      hasBoundaryRight,
      boundaryCount,
      siblingCount
    )
  }

  function isSelected(btn) {
    return btn.type === 'page' && btn.page === value
  }

  function onSelect(btn) {
    if (btn.type === 'page') {
      value = btn.page
    } else {
      value = btn.range[0] + Math.floor((btn.range[1] - btn.range[0]) / 2)
    }
    dispatch('change', { name, type, value })
  }

  function onNav(direction) {
    if (direction === 'prev') {
      value = Math.max(1, value - 1)
    } else {
      value = Math.min(totalPages, value + 1)
    }
    dispatch('change', { name, type, value })
  }
</script>

<div class="tui-pager" data-test-id={testId}>
  <Button
    variant="ghost"
    icon="chevron-left"
    size="small"
    disabled={value === 1}
    on:click={() => onNav('prev')}
  />
  {#each btnData as btn}
    <Button
      variant="ghost"
      size="small"
      selected={isSelected(btn)}
      on:click={() => onSelect(btn)}
    >
      {btn.type === 'page' ? btn.page : '...'}</Button
    >
  {/each}
  <Button
    variant="ghost"
    icon="chevron-right"
    size="small"
    disabled={value === totalPages}
    on:click={() => onNav('next')}
  />
</div>

<style>
  .tui-pager {
    font-family: var(--font-family);
    box-sizing: var(--box-sizing);

    display: flex;
    gap: 8px;
  }
</style>
