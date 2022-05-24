<script>
    export let transaction;
</script>

<tr>
  <td>{transaction.created_at}</td>
  <td><a href="https://www.whatsonchain.com/tx/{transaction.id}">{transaction.id}</a></td>
  <td>{transaction.api_key}</td>
  <td>{transaction.data_bytes}</td>
</tr>
