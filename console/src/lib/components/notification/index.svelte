<script>
  import Icon from '../icon/index.svelte'

  export let testId = null

  export let status = 'success'
  export let title = ''
  export let message = ''

  let statusCol = ''
  let icon = ''

  $: {
    switch (status) {
      case 'success':
        statusCol = '#6EC492'
        icon = 'check-circle'
        break
      case 'failure':
        statusCol = '#FF344C'
        icon = 'exclamation'
        break
    }
  }

  let optProps = {}
  $: {
    if (testId) {
      optProps['data-test-id'] = testId
    }
  }
</script>

<div
  class="tui-notification"
  style:--status-col-local={statusCol}
  {...optProps}
>
  <div class="tab"><Icon name={icon} size={24} /></div>
  <div class="body">
    {#if title}
      <div class="title">{title}</div>
    {/if}
    {#if message}
      <div class="message">{message}</div>
    {/if}
  </div>
  <!-- <div class="close"><Icon name="close" size={24} /></div> -->
</div>

<style>
  .tui-notification {
    font-family: var(--font-family);
    box-sizing: var(--box-sizing);

    position: relative;
    display: flex;
    width: 340px;

    background-color: #ffffff;
    border-radius: 4px;

    /* filter: drop-shadow(0px 0px 12px rgba(35, 45, 124, 0.08)); */
  }

  .tab {
    position: absolute;

    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

    flex: 0 0 auto;
    width: 40px;
    height: 100%;

    background-color: var(--status-col-local);
    border-radius: 4px 0 0 4px;
    color: #ffffff;
  }

  .body {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;

    margin-left: 40px;
    height: 100%;
    padding: 24px 48px 24px 8px;
  }
  .title {
    font-weight: 500;
    font-size: 16px;
    line-height: 26px;
    letter-spacing: -0.02em;

    color: var(--status-col-local);
  }
  .message {
    font-weight: 500;
    font-size: 14px;
    line-height: 24px;
    letter-spacing: -0.02em;

    color: #282933;
  }

  /* .close {
    position: absolute;
    top: 11px;
    right: 11px;
    cursor: pointer;
  } */
</style>
