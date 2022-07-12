<script>
  import { createEventDispatcher } from 'svelte'
  import { headerHeight } from '../../stores'
  import { query } from '../../actions'
  import { medium, large } from '../../styles/breakpoints'

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

  $: isMedium = query(medium)
  $: isLarge = query(large)
  $: mediaSize = $isLarge ? 'large' : $isMedium ? 'medium' : 'small'

  $: {
    gutter = 22
    if (mediaSize === 'large') {
      gutter = 180
    }

    height = 80
    showMobile = false
    showMenu = true
    contentMarginLeft = 0

    if (mediaSize === 'small') {
      height = 60
      showMobile = true
      showMenu = false
      contentMarginLeft = showMobile && !hasMenu ? 60 : 0
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
</script>

<div
  class="tui-header"
  style:--gutter-local={gutter + 'px'}
  style:--height-local={height + 'px'}
  style:--content-margin-left={contentMarginLeft + 'px'}
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
    margin-right: 120px;
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
