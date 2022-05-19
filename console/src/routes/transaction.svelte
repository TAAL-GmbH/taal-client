<script>
    export let transaction;
</script>

<article>
    <table>
        <tr>
          <th>Transaction ID:  &nbsp;</th>
          <th><a href="https://www.whatsonchain.com/tx/{transaction.id}">{transaction.id}</a></th>
        </tr>
      </table>
</article>
<style>
    article {
      margin: 0 0 1em 0;
    }
    h1 {
      font-size: 1.4em;
      margin: 0;
      display: block;
    }
</style>
