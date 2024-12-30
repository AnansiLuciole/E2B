import { CodeInterpreter } from '@e2b/code-interpreter'

const sandbox = await CodeInterpreter.create({ apiKey: 'e2b_b007d9b60ef6b76ce6a6a977875fb480e79e777d' })
await sandbox.notebook.execCell('x = 1')

const execution = await sandbox.notebook.execCell('x+=1; x')
console.log(execution.text)  // outputs 2

await sandbox.close()