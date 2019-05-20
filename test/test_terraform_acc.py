import shlex
import logging
import subprocess
logging.basicConfig(filename='terraform_tests.log', level=logging.INFO)
LOG = logging.getLogger(__name__)


def test_terraform_acc():
    cmd = "make testacc"
    LOG.info("executing command %s ", cmd)
    out = subprocess.check_output(shlex.split(cmd))
    LOG.info("terraform testacc : out %s ", out)
