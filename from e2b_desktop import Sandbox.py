from e2b_desktop import Sandbox
from pkg_resources import Environment

desktop = Sandbox()

from e2b_code_interpreter import CodeInterpreter

with CodeInterpreter(
    api_key="e2b_b007d9b60ef6b76ce6a6a977875fb480e79e777d"
) as code_interpreter:
    sandbox.notebook.exec_cell("x = 1")
    execution = sandbox.notebook.exec_cell("x+=1; x")
    print(execution.text)  # outputs 2
api_key = "e2b_b007d9b60ef6b76ce6a6a977875fb480e79e777d"
set the Environment variable E2B_API_KEY="e2b_b007d9b60ef6b76ce6a6a977875fb480e79e777d"
