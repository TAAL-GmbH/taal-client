<script>
  //
</script>

<div class="tui-popup">
  <div class="title">
    <slot name="title" />
  </div>
  <div class="body">
    <slot name="body" />
  </div>
  <div class="footer">
    <slot name="footer" />
  </div>
</div>

<style>
  .tui-popup {
    width: 100%;
  }

  .title {
    height: 52px;
    width: 100%;
    border-bottom: 1px solid #efefef;
  }
</style>
