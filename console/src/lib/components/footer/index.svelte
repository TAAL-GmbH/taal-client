<script>
  import i18n from '../../i18n'

  $: t = $i18n.t
  const tKey = 'comp.footer'

  export let testId = null

  let year = new Date().getFullYear()
</script>

<div class="tui-footer" data-test-id={testId}>
  <span>{t(`${tKey}.copyright`, { year })}</span>
</div>

<style>
  .tui-footer {
    display: flex;
    align-items: flex-start;
    justify-content: center;

    width: 100%;
    height: 50px;
    flex: 0;

    font-family: 'Work Sans';
    font-style: normal;
    font-weight: 500;
    font-size: 11px;
    line-height: 15px;
    letter-spacing: 0.005em;

    color: #8f8d94;
  }

  span {
    margin-top: 4px;
  }
</style>
