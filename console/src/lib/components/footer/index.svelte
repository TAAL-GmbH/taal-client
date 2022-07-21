<script>
  export let testId = null

  let year = new Date().getFullYear()

  let optProps = {}
  $: {
    if (testId) {
      optProps['data-test-id'] = testId
    }
  }
</script>

<div class="tui-footer" {...optProps}>
  <span>© {year} TAAL TECHNOLOGIES SEZC</span>
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
