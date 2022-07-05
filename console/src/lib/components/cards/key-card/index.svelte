<script>
  import Button from '../../button/index.svelte'
  import Field from '../../field/index.svelte'

  export let key

  export let title = ''
  export let action = 'Add'
  export let usedData = -1
  export let totalData = -1

  function onAction() {
    console.log('onAction: ', action)
  }
</script>

<div class="tui-key-card">
  <div class="header">
    <!-- <div>{title}</div>-->
    <Button variant="secondary" size="small" on:click={onAction}>
      {action}
    </Button>
    <Button variant="secondary" size="small" on:click={onAction}>Remove</Button>
  </div>
  <div class="body">
    <!-- <Field
      label="Available data: 750Mb / 1Gb"
      value={0.75}
      format="progress"
      layout="row"
    /> -->
    <Field
      label="Api key"
      value={key.api_key}
      copy={true}
    />
    <Field
      label="Address"
      value={key.address}
      copy={true}
    />
    <Field label="Date created" value={key.createdAt} />
  </div>
</div>

<style>
  .tui-key-card {
    font-family: var(--font-family);
    box-sizing: var(--box-sizing);

    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }

  .header {
    box-sizing: var(--box-sizing);

    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-end;
    gap: 10px;

    width: 100%;
    height: 60px;
    padding: 0 24px;

    font-weight: 500;
    font-size: 24px;
    line-height: 16px;

    color: #232d7c;
    background-color: #ffffff;
    border-radius: 4px 4px 0px 0px;
  }

  .body {
    box-sizing: var(--box-sizing);

    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 24px;

    width: 100%;
    padding: 15px 24px;

    background: #fcfcff;
    border-radius: 0px 0px 4px 4px;
  }
</style>
