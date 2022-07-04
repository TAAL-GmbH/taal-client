<script>
  import Checkbox from '../../lib/components/checkbox/index.svelte'
  import Heading from '../../lib/components/heading/index.svelte'
  import PageWithMenu from '../../lib/components/page/template/menu/index.svelte'
  import Radio from '../../lib/components/radio/index.svelte'
  import Spacer from '../../lib/components/layout/spacer/index.svelte'
  import Text from '../../lib/components/text/index.svelte'
  import TextInput from '../../lib/components/text-input/index.svelte'

  let textInputs = {
    listenAddress: '',
    taalUrl: '',
    taalTimeout: '',
    fileName: '',
  }
  let checks = {
    checkServer: false,
    checkTransactions: false,
  }
  let radios = {
    dbMode: '',
  }

  function onChange(e) {
    const { name, group, type, value, checked } = e.detail
    // console.log(
    //   'name =',
    //   radios,
    //   ' group =',
    //   group,
    //   ' type =',
    //   type,
    //   ' checked =',
    //   checked
    // )
    switch (type) {
      case 'text':
        textInputs[name] = value || ''
        break
      case 'checkbox':
        checks[name] = checked || false
        break
      case 'radio':
        radios[group] = checked ? name : ''
        radios = radios
        break
    }
  }

  function onInputMount(e) {
    e.detail.inputRef.focus()
  }
</script>

<PageWithMenu>
  <div class="island">
    <Heading value="Settings" />
    <Spacer h={24} />
    <div class="sub-row">
      <Heading value="Server settings" size={2} />
    </div>
    <Spacer h={24} />
    <TextInput
      name="listenAddress"
      label="Listen address"
      value={textInputs['listenAddress']}
      on:change={onChange}
      on:mount={onInputMount}
    />
    <Spacer h={8} />
    <Text
      size={5}
      color="#8F8D94"
      value={'TaalClient will listen on the specified address. The default is "localhost:9500" To bind to all interfaces on your machine, omit the host part and only specify the port, for example ":9500"'}
    />
    <Spacer h={24} />
    <TextInput
      name="taalUrl"
      label="TAAL URL"
      value={textInputs['taalUrl']}
      on:change={onChange}
    />
    <Spacer h={8} />
    <Text
      size={5}
      color="#8F8D94"
      value={'TaalClient communicates with Taal API servers which are hosted at https://api.taal.com.'}
    />
    <Spacer h={24} />
    <TextInput
      name="taalTimeout"
      label="TAAL timeout"
      value={textInputs['taalTimeout']}
      on:change={onChange}
    />
    <Spacer h={8} />
    <Text
      size={5}
      color="#8F8D94"
      value={'The default timeout is 10 seconds. This settings can be provided in milliseconds (ms), seconds (s) or minutes (m).'}
    />
    <Spacer h={24} />
    <Text size={4} value={'Debug'} />
    <Spacer h={8} />
    <Text
      size={5}
      color="#8F8D94"
      value={'Debug modes will add extra output to the console output of this service. Server debug will output all access to server endpoints. Transactions debug will output funding and data transaction raw hex.'}
    />
    <Spacer h={16} />
    <Checkbox
      name="checkServer"
      label="Server"
      checked={checks['checkServer']}
      labelPlacement="right"
      on:change={onChange}
    />
    <Spacer h={16} />
    <Checkbox
      name="checkTransactions"
      label="Transactions"
      checked={checks['checkTransactions']}
      labelPlacement="right"
      on:change={onChange}
    />
    <Spacer h={32} />
    <div class="sub-row">
      <Heading value="Transaction data" size={2} />
    </div>
    <Spacer h={24} />
    <Text size={4} value={'Database mode'} />
    <Spacer h={16} />
    <div class="radio-row">
      <Radio
        name="radioLocal"
        group="dbMode"
        label="Local"
        checked={radios['dbMode'] === 'radioLocal'}
        labelPlacement="right"
        on:change={onChange}
      />
      <Radio
        name="radioRemote"
        group="dbMode"
        label="Remote"
        checked={radios['dbMode'] === 'radioRemote'}
        labelPlacement="right"
        on:change={onChange}
      />
    </div>
    <Spacer h={24} />
    <TextInput
      name="fileName"
      label="File name"
      value={textInputs['fileName']}
      on:change={onChange}
    />
  </div>
</PageWithMenu>

<style>
  .island {
    display: flex;
    flex-direction: column;
    width: 100%;
    max-width: 920px;
    padding-top: 40px;
    margin-bottom: 100px;
  }

  .sub-row {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .radio-row {
    display: flex;
    align-items: center;
    gap: 24px;
  }
</style>
