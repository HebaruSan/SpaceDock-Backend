from SpaceDock.app import app
from SpaceDock.config import cfg

# Start up the app
if __name__ == '__main__':
    app.run(host = cfg['host'], port = cfg.geti('port'), threaded = cfg.getb('threaded'))