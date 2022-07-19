<script>
  import { createEventDispatcher } from 'svelte'
  import Button from '../button/index.svelte'
  import { getInputLabel } from '../../utils/strings'

  const dispatch = createEventDispatcher()

  export let label = ''
  export let labelPlacement = 'top'
  export let labelAlignment = 'start' // 'start' | 'center' | 'end'
  export let required = false
  export let name = ''
  export let value = ''
  // export let multiple = false
  export let items = []
  export let disabled = false
  export let valid = true
  export let error = ''
  export let maxVisibleListItems = 6

  export let toggle = true
  export let focusRect = true
  export let padding = '6px'
  export let gap = '4px'

  let type = 'select'

  // direction
  let direction = 'row'

  $: {
    switch (labelPlacement) {
      case 'top':
        direction = 'column'
        break
      case 'bottom':
        direction = 'column-reverse'
        break
      case 'left':
        direction = 'row'
        break
      case 'right':
        direction = 'row-reverse'
        break
    }
  }

  let labelAlign = 'center'

  $: {
    switch (labelAlignment) {
      case 'start':
        labelAlign = 'flex-start'
        break
      case 'center':
        labelAlign = 'center'
        break
      case 'end':
        labelAlign = 'flex-end'
        break
    }
  }

  // size
  export let size = 'small'

  let height = 60
  let fontSize = '16px'
  let listHeight = 0

  $: {
    switch (size) {
      case 'large':
        height = 60
        fontSize = '16px'
        break
      case 'medium':
        height = 48
        fontSize = '16px'
        break
      case 'small':
        height = 36
        fontSize = '14px'
        break
    }

    listHeight = items.length * height
    if (maxVisibleListItems !== -1) {
      listHeight = Math.min(listHeight, maxVisibleListItems * height)
    }
    listHeight += 8
  }

  let focused = false

  function onItemSelect(val) {
    value = val
    dispatch('change', { name, type, value })
  }
</script>

<div
  class="tui-button-select"
  tabindex={0}
  style:--height-local={height + 'px'}
  style:--padding-local={padding}
  style:--gap-local={gap}
  style:--fontSize-local={fontSize}
  style:--label-align-local={labelAlign}
  style:--direction-local={direction}
  style:--list-height-local={listHeight + 'px'}
  on:focus={(e) => (focused = true)}
  on:blur={(e) => (focused = false)}
>
  <div class="placement">
    {#if label}
      <label for={name}>{getInputLabel(label, required)}</label>
    {/if}
    <div
      class="button-select"
      class:disabled
      class:error={!valid || error !== ''}
      class:focused={focusRect && focused}
    >
      {#each items as item (item.value)}
        <Button
          variant="ghost"
          selected={toggle ? item.value === value : false}
          icon={item.icon}
          iconAfter={item.iconAfter}
          {disabled}
          {size}
          toggle={true}
          on:click={() => onItemSelect(item.value)}
        >
          {item.label}
        </Button>
      {/each}
    </div>
  </div>
  {#if error}
    <div class="error-msg">{error}</div>
  {/if}
</div>

<style>
  .tui-button-select {
    font-family: var(--font-family);
    box-sizing: var(--box-sizing);
    outline: none;

    display: flex;
    flex-direction: column;
    flex-grow: 0;
    align-self: flex-start;
  }

  .placement {
    display: flex;
    flex-direction: var(--direction-local);
    align-items: var(--label-align-local);
    gap: 8px;
  }

  label {
    font-weight: 500;
    font-size: 14px;
    line-height: 16px;

    color: #232d7c;
  }

  .button-select {
    display: flex;
    align-items: center;
    gap: var(--gap-local);
    padding: var(--padding-local);
    height: var(--height-local);

    background-color: #ffffff;
    border-style: solid;
    border-width: 1px;
    border-color: #ffffff;
    border-radius: 4px;
  }
  .button-select.disabled {
    cursor: auto;
  }

  .button-select.focused {
    border-width: 2px;
    border-color: #4a3aff;
    padding: max(calc(var(--padding-local) - 1px), 0px);
  }

  .button-select.error,
  .button-select.error.focused {
    border-width: 1px;
    border-color: #ff344c;
    padding: var(--padding-local);
  }

  .error-msg {
    margin-top: 8px;

    font-weight: 500;
    font-size: 12px;
    line-height: 16px;
    letter-spacing: -0.01em;

    color: #ff344c;
  }

  .disabled,
  .disabled:active,
  .disabled.focused {
    padding: var(--padding-local);
    border-width: 1px;
    border-color: #8f8d94;
  }
</style>
