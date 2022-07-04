<script>
  import { createEventDispatcher } from 'svelte'

  import Button from '../button/index.svelte'
  import Icon from '../icon/index.svelte'
  import Logo from '../logo/index.svelte'

  const dispatch = createEventDispatcher()

  export let links = []
  export let actions = []

  export let showMobile = false
  export let open = true
  export let showLinks = true
  export let showActions = true
  export let dataKey = 'path'

  function onLink(item) {
    dispatch('link', item)
  }

  function onAction(item) {
    dispatch('action', item)
  }

  function onToggle() {
    dispatch('toggle-menu')
  }
</script>

<div class="tui-header">
  {#if showMobile}
    <div class="icon" on:click={(e) => onToggle()}>
      <Icon name={open ? 'close' : 'menu'} />
    </div>
  {/if}
  <div class="content">
    <div class="logo">
      <Logo name="taal-blue" height={38} width={93} />
      <Logo name="client" height={38} width={116} />
    </div>
    {#if showLinks}
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
    display: flex;
    align-items: center;

    position: absolute;
    top: 0;
    right: 0;
    left: 0;

    height: 80px;
    max-height: 80px;

    background-color: #ffffff;
  }

  .icon {
    position: absolute;
    left: 15px;
    top: 28px;
    cursor: pointer;
    z-index: 1;
  }

  .content {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: stretch;

    width: calc(100% - 120px);
    max-width: calc(100% - 120px);
    padding: 0 60px;

    color: blue;
  }

  .logo {
    display: flex;
    align-items: center;
    margin-right: 130px;
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
