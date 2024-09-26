const core = require("@actions/core");
const gh = require("@actions/github");

// Get input
try {

	// Will appear if `ACTIONS_STEP_DEBUG=true`
	core.debug("Not too important message")
	core.error("Will show up red")
	core.warning("Will show up yellow")

	// obfuscate
	core.setSecret("secret")

	// logging group
	core.startGroup("INPUT")
	core.startGroup("Getting input...")
	const name = core.getInput("who-to-greet");
	console.log(`Hello ${name}`);
	core.startGroup("Got output")
	core.endGroup()

	// export env var
	// can be used in follow up command to this action
	// run: |
	//   echo "Time: ${{ steps.greet.outputs.time }}"
	//   echo $KEY
	core.exportVariable("KEY", "VALUE")

	
	// Set output
	const time = new Date();
	core.setOutput("time", time.toTimeString())
	
	
	// Get GitHub metadata
	console.log(JSON.stringify(gh, null, 3))
} catch(error) {
	core.setFailed(`Why it failed: ${error.message}`)
}



// To fail the action, we use
// core.setFailed("Why it failed");