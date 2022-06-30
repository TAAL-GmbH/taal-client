<script>
  import Icon from '../icon/index.svelte'
  import Progress from '../progress/index.svelte'

  export let label = ''
  export let value = null
  export let format = null
  export let layout = 'column'
  export let copy = false

  function onCopy(value) {
    // TODO: copy to clipboard
    console.log('onCopy: value = ', value)
  }
</script>

<div class="tui-field" class:row={layout === 'row'}>
  <div class="label">{label}</div>
  <div class="value">
    <div>
      {#if format === 'progress'}
        <div style="width: 100px;">
          <Progress ratio={value} />
        </div>
      {:else}
        {value}
      {/if}
    </div>
    {#if copy}
      <div class="ico" on:click={() => onCopy(value)}>
        <Icon name="document-duplicate" size={18} />
      </div>
    {/if}
  </div>
</div>

<style>
  .tui-field {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }
  .tui-field.row {
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    width: 100%;
  }

  .label {
    font-weight: 600;
    font-size: 13px;
    line-height: 16px;

    text-transform: uppercase;

    color: #232d7c;
  }

  .value {
    display: flex;
    align-items: center;
    gap: 8px;

    font-weight: 400;
    font-size: 14px;
    line-height: 24px;
    letter-spacing: -0.01em;

    color: #282933;
  }

  .ico {
    width: 18px;
    height: 18px;
    cursor: pointer;
  }
</style>
