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
    arrowFocusIndex = -1
    dispatch('change', { name, type, value })
  }

  let arrowFocusIndex = -1

  function onKeyDown(e) {
    if (!e) e = window.event
    const keyCode = e.code || e.key
    switch (keyCode) {
      case 'ArrowRight':
      case 'ArrowLeft':
        e.preventDefault()
        if (arrowFocusIndex === -1) {
          initArrowFocusIndex()
        }
        arrowFocusIndex =
          keyCode === 'ArrowRight'
            ? (arrowFocusIndex + 1) % items.length
            : arrowFocusIndex === 0
            ? items.length - 1
            : (arrowFocusIndex - 1) % items.length
        return false
      case 'Enter':
        e.preventDefault()
        if (arrowFocusIndex !== -1) {
          document
            .querySelectorAll('.button-select > *')
            [arrowFocusIndex].click()
          arrowFocusIndex = -1
        }
        return false
    }
  }

  function initArrowFocusIndex() {
    const result = items.filter((item) => item.value === value)
    if (result && result.length > 0) {
      arrowFocusIndex = items.indexOf(result[0])
    } else {
      arrowFocusIndex = 0
    }
  }

  function onFocusAction(eventName) {
    switch (eventName) {
      case 'blur':
        focused = false
        break
      case 'focus':
        focused = true
        arrowFocusIndex = -1
        break
    }
    dispatch(eventName)
  }
</script>

<div
  class="tui-button-select"
  style:--height-local={height + 'px'}
  style:--padding-local={padding}
  style:--gap-local={gap}
  style:--fontSize-local={fontSize}
  style:--label-align-local={labelAlign}
  style:--direction-local={direction}
  style:--list-height-local={listHeight + 'px'}
  tabindex={-1}
>
  <div class="placement">
    {#if label}
      <label for={name}>{getInputLabel(label, required)}</label>
    {/if}
    <div
      class="button-select"
      tabindex={0}
      class:disabled
      class:error={!valid || error !== ''}
      class:focused={focusRect && focused}
      on:focus={() => onFocusAction('focus')}
      on:blur={() => onFocusAction('blur')}
      on:keydown={onKeyDown}
    >
      {#each items as item, i (item.value)}
        <Button
          variant="ghost"
          selected={toggle ? item.value === value : false}
          icon={item.icon}
          iconAfter={item.iconAfter}
          {disabled}
          {size}
          toggle={true}
          class={focused && arrowFocusIndex === i ? 'hover' : ''}
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

    outline: none;
  }
  .button-select.disabled {
    cursor: auto;
  }

  .button-select.focused {
    border-width: 1px;
    border-color: #4a3aff;
    padding: var(--padding-local);
  }
  .button-select.focused:focus-visible {
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
