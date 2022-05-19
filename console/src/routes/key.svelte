<script>
    export let key;
</script>

<article>
    <table>
      <tr>
        <th>API key:  &nbsp;</th>
        <th>{key.api_key}</th>
      </tr>
      <tr>
        <th>Public key:  &nbsp;</th>
        <th>{key.public_key}</th>
      </tr>
      <tr>
        <th>Address:  &nbsp;</th>
        <th>{key.address}</th>
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
