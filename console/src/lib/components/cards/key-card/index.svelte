<script>
  import Button from '../../button/index.svelte'
  import Field from '../../field/index.svelte'

  export let title = ''
  export let action = 'Upgrade'
  export let usedData = -1
  export let totalData = -1

  function onAction() {
    console.log('onAction: ', action)
  }
</script>

<div class="tui-key-card">
  <div class="header">
    <div>{title}</div>
    <Button variant="secondary" size="small" on:click={onAction}
      >{action}</Button
    >
  </div>
  <div class="body">
    <Field
      label="Available data: 750Mb / 1Gb"
      value={0.75}
      format="progress"
      layout="row"
    />
    <Field
      label="Api key"
      value="mainnet_1PyWzkfKrq1kakvLTeaCdAL8y8UJAcZAqU"
      copy={true}
    />
    <Field
      label="Address"
      value="1L7GLWPyWzkfKrq1kakvLTeaCdAL8y8UJAcZAqU"
      copy={true}
    />
    <Field label="Date created" value="2022-06-08 12:32:03" />
  </div>
</div>

<style>
  .tui-key-card {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }

  .header {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;

    width: calc(100% - 48px);
    height: 60px;
    padding: 15px 24px;

    font-weight: 500;
    font-size: 24px;
    line-height: 16px;

    color: #232d7c;
    background-color: #ffffff;
    border-radius: 4px 4px 0px 0px;
  }

  .body {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 24px;

    width: calc(100% - 48px);
    padding: 15px 24px;

    background: #fcfcff;
    border-radius: 0px 0px 4px 4px;
  }
</style>
