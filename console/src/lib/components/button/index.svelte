<script>
  import Icon from '../icon/index.svelte'

  // variant
  export let variant = 'primary'

  // icon
  export let icon = null
  export let iconAfter = null

  let hasIcon = false
  let direction = 'row'

  $: {
    hasIcon = icon || iconAfter

    if (hasIcon) {
      direction = icon ? 'row' : 'row-reverse'
    }
  }

  let primary = false
  let secondary = false
  let ghost = false
  let link = false

  $: {
    primary = variant === 'primary'
    secondary =
      variant === 'secondary' || variant === 'ghost' || variant === 'link'
    ghost = variant === 'ghost'
    link = variant === 'link'
  }

  export let disabled = false
  export let selected = false
  export let toggle = false
  export let width = -1

  // size
  export let size = 'medium'

  let height = '60px'
  let padding = '10px 32px'
  let fontSize = '16px'
  let iconSize = 24

  $: {
    switch (size) {
      case 'large':
        height = '60px'
        padding = '10px 32px'
        fontSize = '16px'
        break
      case 'medium':
        height = '48px'
        padding = '10px 32px'
        fontSize = '16px'
        break
      case 'small':
        height = '36px'
        padding = '4px 16px'
        fontSize = '14px'
        iconSize = 18
        break
    }

    if (link) {
      padding = '0'
      fontSize = '16px'
    }
  }
</script>

<div
  class="tui-button"
  class:primary
  class:secondary
  class:ghost
  class:link
  class:disabled
  class:selected
  class:toggle
  style:--height-local={height}
  style:--padding-local={padding}
  style:--fontSize-local={fontSize}
  style:--direction-local={direction}
  style:--width-local={width === -1 ? 'auto' : `${width}px`}
  on:click
>
  {#if hasIcon}
    <div class="icon" style="width: {iconSize}px; height: {iconSize}px;">
      <Icon name={icon || iconAfter} size={iconSize} />
    </div>
  {/if}
  {#if $$slots.default}
    <div class="label"><slot /></div>
  {/if}
</div>

<style>
  .tui-button {
    display: flex;
    flex-direction: var(--direction-local);
    align-items: center;
    justify-content: center;
    gap: 10px;

    outline: none;
    box-sizing: var(--box-sizing);

    width: var(--width-local);
    min-height: var(--height-local);
    max-height: var(--height-local);
    padding: var(--padding-local) !important;

    font-family: var(--font-family);
    font-size: var(--fontSize-local);
    font-weight: 500;
    line-height: 30px;
    letter-spacing: -0.01em;

    border-width: 1px;
    border-style: solid;
    border-radius: 4px;
    cursor: pointer;
    user-select: none;
  }

  .icon {
    margin: 0 -6px;
  }

  .tui-button .label {
    display: flex;
    align-items: center;
    white-space: nowrap;
  }

  .primary {
    background-color: #232d7c;
    border-color: #232d7c;
    color: #ffffff;
  }
  .primary:hover {
    background-color: #1e276c;
  }
  .primary:focus {
    background-color: #1e276c;
    border-color: #4a3aff;
  }
  .primary:active {
    background-color: #181f54;
  }

  .secondary {
    background-color: #ffffff;
    border-color: #232d7c;
    color: #232d7c;
  }
  .secondary:hover {
    background-color: #f4f6ff;
  }
  .secondary:focus {
    background-color: #f4f6ff;
    border-color: #4a3aff;
  }
  .secondary:active {
    background-color: #dbdbff;
  }

  .ghost {
    border: none;
  }
  .ghost.disabled.toggle {
    background-color: transparent;
  }
  .ghost.selected {
    background-color: #dbdbff;
  }
  .ghost.selected.disabled {
    background-color: #efefef;
  }

  .link {
    background-color: transparent !important;
    border-radius: 0;
    border: none;
    border-top: 2px solid transparent;
    border-bottom: 2px solid transparent;
  }
  .link:hover,
  .link.selected {
    border-bottom: 2px solid #232d7c;
  }
  .link.selected {
    font-weight: 600;
  }

  .disabled,
  .disabled:hover {
    background-color: #efefef;
    border-color: #efefef;
    color: #8f8d94;

    cursor: auto;
    pointer-events: none;
  }
</style>
