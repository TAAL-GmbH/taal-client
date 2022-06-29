<script>
  import { createEventDispatcher } from 'svelte'
  import { fade, slide } from 'svelte/transition'
  import { clickOutside } from '../../actions'
  import Icon from '../icon/index.svelte'
  import { getInputLabel } from '../../utils/strings'

  const dispatch = createEventDispatcher()

  export let label = ''
  export let labelPlacement = 'top'
  export let required = false
  export let name = ''
  export let value = 0
  export let defaultValue = ''
  export let multiple = false
  export let items = []
  export let disabled = false
  export let error = ''
  export let maxVisibleListItems = 6

  let open = false

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

  // size
  export let size = 'medium'

  let height = 60
  let padding = '10px 32px'
  let fontSize = '16px'
  let iconSize = 16
  let listHeight = 0

  $: {
    switch (size) {
      case 'large':
        height = 60
        padding = '10px 32px'
        fontSize = '16px'
        break
      case 'medium':
        height = 48
        padding = '10px 32px'
        fontSize = '16px'
        break
      case 'small':
        height = 36
        padding = '4px 16px'
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

  let selectRef

  function onSelectParentClick() {
    selectRef.focus()
    open = !open
  }

  function onSelectChange(e) {
    dispatch('change', { name, value: e.data })
  }

  function onClose() {
    open = false
  }

  function onItemSelect(val) {
    value = val
    open = false
    dispatch('change', { name, value })
  }
</script>

<div
  class="tui-dropdown"
  class:open
  style:--height-local={height + 'px'}
  style:--padding-local={padding}
  style:--fontSize-local={fontSize}
  style:--direction-local={direction}
  style:--icon-size-local={iconSize + 'px'}
  style:--list-height-local={listHeight + 'px'}
  use:clickOutside
  on:outclick={onClose}
  on:focus={(e) => (focused = true)}
>
  <div class="placement">
    {#if label}
      <label for={name}>{getInputLabel(label, required)}</label>
    {/if}
    <div
      class="select"
      class:disabled
      class:error={error !== ''}
      class:focused
      on:click={disabled ? null : onSelectParentClick}
    >
      <select
        bind:this={selectRef}
        {multiple}
        {defaultValue}
        {disabled}
        {name}
        {value}
        on:change={onSelectChange}
        on:focus={(e) => (focused = true)}
        on:blur={(e) => (focused = false)}
      >
        {#each items as item (item.value)}
          <option value={item.value}>
            {item.label}
          </option>
        {/each}
      </select>
      <div class="icon">
        <Icon name="chevron-down" size={iconSize} />
      </div>
    </div>
    {#if open}
      <div class="list-container">
        <div in:slide out:fade class="list">
          {#each items as item (item.value)}
            <div
              class="list-item"
              class:selected={item.value === value}
              on:click={() => onItemSelect(item.value)}
            >
              {item.label}
            </div>
          {/each}
        </div>
      </div>
    {/if}
  </div>
  {#if error}
    <div class="error-msg">{error}</div>
  {/if}
</div>

<style>
  .tui-dropdown {
    font-family: var(--font-family);
    box-sizing: var(--box-sizing);

    display: flex;
    flex-direction: column;
    flex-grow: 1;
  }

  .icon {
    width: var(--icon-size-local);
    height: var(--icon-size-local);
    transition: transform 0.2s ease-in-out;
  }
  .tui-dropdown.open .icon {
    transform: rotate(0deg);
  }
  .tui-dropdown.open .icon {
    transform: rotate(180deg);
  }

  .list-container {
    position: relative;
    z-index: 200;
    margin-top: -8px;
  }
  .list {
    position: absolute;
    top: -1px;
    width: 100%;
    height: var(--list-height-local);
    overflow-x: hidden;
    overflow-y: auto;

    box-sizing: var(--box-sizing);
    background-color: #ffffff;
    border-style: solid;
    border-width: 1px;
    border-color: #232d7c;
    border-radius: 0 0 4px 4px;

    border-top: 1px dashed #8f8d94;
  }
  .list-item {
    display: flex;
    align-items: center;
    width: 100%;
    height: var(--height-local);

    padding: 1px 18px;
    background: white;

    font-weight: 400;
    font-size: 14px;
    line-height: 24px;
    letter-spacing: -0.01em;

    color: #282933;
    cursor: pointer;
  }
  .list-item:hover {
    background: #f4f6ff;
  }
  .list-item.selected {
    background: #dbdbff;
  }

  .placement {
    display: flex;
    flex-direction: var(--direction-local);
    gap: 8px;
  }

  label {
    font-weight: 500;
    font-size: 14px;
    line-height: 16px;

    color: #232d7c;
  }

  select {
    outline: none;
    border: none;
    width: calc(100% - var(--icon-size-local));
    background: none;
    appearance: none;
    pointer-events: none;
  }

  .select {
    display: flex;
    align-items: center;
    padding: 1px 16px;
    height: var(--height-local);

    background-color: #ffffff;
    border-style: solid;
    border-width: 1px;
    border-color: #232d7c;
    border-radius: 4px;

    font-weight: 400;
    font-size: 16px;
    line-height: 24px;
    letter-spacing: -0.01em;

    color: #282933;
    cursor: pointer;
  }
  .tui-dropdown.open .select {
    border-radius: 4px 4px 0 0;
  }
  .select.disabled {
    cursor: auto;
  }

  .select.focused {
    border-width: 2px;
    border-color: #4a3aff;
    padding: 0px 15px;
  }
  .tui-dropdown.open .select.focused {
    padding: 1px 16px;
    border-width: 1px;
    border-color: #232d7c;
  }

  .select.error,
  .select.error.focused {
    border-width: 1px;
    border-color: #ff344c;
    padding: 1px 16px;
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
  .disabled:active {
    background-color: #efefef;
    border-color: #232d7c;
  }
</style>
