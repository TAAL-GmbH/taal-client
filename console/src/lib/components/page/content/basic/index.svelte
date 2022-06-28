<div class="page-content">
  <slot />
</div>

<style>
  .page-content {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-around;

    width: 100%;
    min-height: calc(100% - 50px);

    width: calc(100% - 120px);
    max-width: calc(100% - 120px);
    padding: 0 60px;

    overflow: hidden;
  }
</style>
