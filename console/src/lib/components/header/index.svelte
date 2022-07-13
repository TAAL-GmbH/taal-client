<script>
  import { createEventDispatcher } from 'svelte'
  import { headerHeight, mediaSize } from '../../stores'

  import Button from '../button/index.svelte'
  import Icon from '../icon/index.svelte'
  import Logo from '../logo/index.svelte'

  const dispatch = createEventDispatcher()

  export let links = []
  export let actions = []

  $: hasMenu = links.length > 0

  let open = false
  export let showLinks = true
  export let showActions = true
  export let dataKey = 'path'

  let gutter = 22
  let height = 80
  let showMobile = false
  let showMenu = true
  let contentMarginLeft = 0
  let logoMarginRight = 120

  $: {
    gutter = 22
    if ($mediaSize === 'large') {
      gutter = 180
    }

    height = 80
    showMobile = false
    showMenu = true
    contentMarginLeft = 0
    logoMarginRight = 120

    if ($mediaSize === 'small') {
      height = 60
      showMobile = true
      showMenu = false
      contentMarginLeft = showMobile && !hasMenu ? 60 : 0
      logoMarginRight = 10
    }

    $headerHeight = height
  }

  function onLink(item) {
    dispatch('link', item)
  }

  function onAction(item) {
    dispatch('action', item)
  }

  function onToggle() {
    open = !open
    dispatch('toggle-menu', { open })
  }

  // box-shadow: 0px 1px 10px rgba(46, 46, 46, 0.1);
</script>

<div
  class="tui-header"
  class:small={$mediaSize === 'small'}
  style:--gutter-local={gutter + 'px'}
  style:--height-local={height + 'px'}
  style:--content-margin-left={contentMarginLeft + 'px'}
  style:--logo-margin-right={logoMarginRight + 'px'}
>
  {#if showMobile && hasMenu}
    <div class="icon" on:click={(e) => onToggle()}>
      <Icon name={open ? 'close' : 'menu'} size={30} />
    </div>
  {/if}
  <div class="content">
    <div class="logo">
      <Logo name="taal-blue" height={38} width={93} />
      <Logo name="client" height={38} width={116} />
    </div>
    {#if showLinks && hasMenu && showMenu}
      <div class="links">
        {#each links as link (link[dataKey])}
          <Button
            variant="link"
            size="small"
            selected={link.selected}
            on:click={(e) => onLink(link)}>{link.label}</Button
          >
        {/each}
      </div>
    {/if}
    <div class="spacer" />
    {#if showActions}
      <div class="actions">
        {#each actions as action (action[dataKey])}
          <Button
            variant={action.selected ? 'primary' : 'secondary'}
            size="small"
            on:click={(e) => onAction(action)}>{action.label}</Button
          >
        {/each}
      </div>
    {/if}
  </div>
</div>

<style>
  .tui-header {
    font-family: var(--font-family);
    box-sizing: var(--box-sizing);
    display: flex;
    align-items: center;

    position: absolute;
    top: 0;
    right: 0;
    left: 0;

    height: var(--height-local);
    max-height: var(--height-local);
    width: 100%;
    padding: 0 var(--gutter-local);

    background-color: #ffffff;
  }
  .tui-header.small {
    box-shadow: 0px 1px 10px rgba(46, 46, 46, 0.1);
  }

  .icon {
    width: 30px;
    height: 30px;
    margin-right: 30px;
    cursor: pointer;
    color: #232d7c;
  }

  .content {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: stretch;

    width: 100%;
    margin-left: var(--content-margin-left);

    color: blue;
  }

  .logo {
    display: flex;
    align-items: center;
    margin-right: var(--logo-margin-right);
  }
  .links {
    display: flex;
    flex: 0;
    gap: 20px;
  }
  .spacer {
    flex: 1;
  }
  .actions {
    align-self: flex-end;
    display: flex;
    align-items: center;
    gap: 10px;
  }
</style>
