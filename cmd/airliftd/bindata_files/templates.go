package bindata_files

import (
	"time"

	"ktkr.us/pkg/vfs/bindata"
)

func init() {
	bindata.RegisterFile("templates/config.tmpl", time.Unix(1442595507, 0), []byte("{{ define \"config\" }}\n{{ with .Data }}\n  <section id=\"section-overview\" class=\"floating-section\">\n    {{ template \"overview\" . }}\n  </section>\n  <section id=\"section-config\" class=\"floating-section\">\n    <h1>Configuration</h1>\n    <form id=\"config\">\n      <div class=\"box\" id=\"host-box\" data-tooltip=\"Returned file links will begin with this domain and path.\" data-tt-pos=\"top\">\n        <label for=\"host\">Base URL</label>\n        <input type=\"text\" id=\"host\" name=\"host\" value=\"{{ .Conf.Host }}\" placeholder=\"i.example.com\">\n      </div>\n      <div class=\"box\" id=\"id-box\">\n        /<span id=\"sample-id\"></span><span id=\"sample-ext\">.ext</span>\n      </div>\n      <div class=\"box\">\n        <label for=\"id-size\">Length of File ID</label>\n        <input type=\"range\" id=\"id-size\" name=\"id-size\" min=\"2\" max=\"12\" value=\"{{ .Conf.HashLen }}\">\n      </div>\n      <div class=\"box checkbox\" data-tooltip=\"Enable to append the original file extension to returned links.\" data-tt-pos=\"left\">\n        <input type=\"checkbox\" id=\"append-ext\" name=\"append-ext\"{{ if .Conf.AppendExt }} checked{{ end }}>\n        <label for=\"append-ext\">Append file extensions</label>\n      </div>\n      <div class=\"box check-enable\">\n        <input type=\"checkbox\" class=\"hider\" id=\"enable-age-prune\" name=\"enable-age-prune\"{{ if .Conf.MaxAgeEnable }} checked{{ end }}>\n        <label for=\"enable-age-prune\">Limit upload age</label>\n        <div class=\"hidee\">\n          <label for=\"max-age\">Maximum age (days)</label>\n          <input type=\"number\" id=\"max-age\" name=\"max-age\" value=\"{{ .Conf.Age }}\" min=\"0\"{{ if not .Conf.MaxAgeEnable }} disabled{{ end }}>\n        </div>\n      </div>\n      <div class=\"box check-enable\">\n        <input type=\"checkbox\" class=\"hider\" id=\"enable-size-prune\" name=\"enable-size-prune\"{{ if .Conf.MaxSizeEnable }} checked{{ end }}>\n        <label for=\"enable-size-prune\">Limit total uploads size</label>\n        <div class=\"hidee\">\n          <label for=\"max-size\">Maximum size (MB)</label>\n          <input type=\"number\" id=\"max-size\" name=\"max-size\" value=\"{{ .Conf.Size }}\" min=\"0\"{{ if not .Conf.MaxSizeEnable }} disabled{{ end }}>\n        </div>\n      </div>\n      <div class=\"box check-enable\" data-tooltip=\"Enable to allow uploads to show Twitter Cards with file previews if applicable.\" data-tt-pos=\"left\">\n        <input type=\"checkbox\" class=\"hider\" name=\"twitter-card\"{{ if .Conf.TwitterCardEnable }} checked{{ end }}>\n        <label for=\"twitter-card\">Enable Twitter Cards</label>\n        <div class=\"hidee\">\n          <label for=\"twitter-handle\">Twitter handle</label>\n          <input type=\"text\" id=\"twitter-handle\" name=\"twitter-handle\" value=\"{{ .Conf.TwitterHandle }}\" required placeholder=\"@handle\"{{ if not .Conf.TwitterCardEnable }} disabled{{ end }}>\n        </div>\n      </div>\n      <div class=\"box\" id=\"directory-box\">\n        <label for=\"directory\">Upload Directory</label>\n        <input type=\"text\" id=\"directory\" name=\"directory\" value=\"{{ .Conf.Directory }}\" placeholder=\"/home/user/uploads\">\n      </div>\n      <div class=\"box\" id=\"newpass-box\" data-tooltip=\"Enter a new password here to change your password.\" data-tt-pos=\"right\">\n        <label for=\"newpass\">New password</label>\n        <input type=\"password\" id=\"newpass\" name=\"newpass\" placeholder=\"\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\xe2\x80\xa2\">\n      </div>\n      <hr>\n      <div class=\"box\" id=\"password-box\" data-tooltip=\"Current password is required to update configuration\" data-tt-pos=\"left\">\n        <label for=\"password\">Current password</label>\n        <input type=\"password\" id=\"password\" name=\"password\" required placeholder=\"Required\">\n      </div>\n      <button id=\"submit\" type=\"button\">Update configuration</button>\n    </form>\n  </section>\n  <script>window.addEventListener('load', configPage, false);</script>\n{{ end }}\n{{ end }}\n\n{{ define \"overview\" }}\n    <h1>Overview</h1>\n    <p><strong><a href=\"/-/history/0\">{{ .NumUploads }} upload{{ if ne .NumUploads 1 }}s{{ end }}</a></strong> totalling <strong>{{ .UploadsSize }}</strong>. (<a href=\"javascript:purgeAll()\">purge</a>)</p>\n    <p>Thumbnail cache is <strong>{{ .ThumbsSize }}</strong>. (<a href=\"javascript:purgeThumbs()\">purge</a>)</p>\n{{ end }}\n"))
	bindata.RegisterFile("templates/errors/errors.tmpl", time.Unix(1440218376, 0), []byte("{{ define \"400\" }}<!doctype html>\n<html>\n  <head>\n    <title>400</title>\n    <link rel=\"stylesheet\" href=\"/static/style.css\">\n  </head>\n  <body>\n    <div class=\"error\">\n      <h1>You're doing it wrong.</h1>\n      {{ if .Err }}<p>{{ .Err }}</p>{{ end }}\n    </div>\n  </body>\n</html>\n{{ end }}\n{{ define \"404\" }}<!doctype html>\n<html>\n  <head>\n    <title>404</title>\n    <link rel=\"stylesheet\" href=\"/static/style.css\">\n  </head>\n  <body>\n    <div class=\"error\">\n      <h1>This isn't the page you're looking for.</h1>\n    </div>\n  </body>\n</html>\n{{ end }}\n{{ define \"500\" }}<!doctype html>\n<html>\n  <head>\n    <title>500</title>\n    <link rel=\"stylesheet\" href=\"/static/style.css\">\n  </head>\n  <body>\n    <div class=\"error\">\n      <h1>Something went wrong.</h1>\n      {{ if .Err }}<p>{{ .Err }}</p>{{ end }}\n    </div>\n  </body>\n</html>\n{{ end }}\n"))
	bindata.RegisterFile("templates/history.tmpl", time.Unix(1442347808, 0), []byte("{{ define \"history\" }}\n{{ with .Data }}\n<section id=\"history\">\n  {{ if len .List | lt 25 }}{{ template \"pagination\" . }}{{ end }}\n  <ul>\n    {{ range .List }}\n    <li class=\"history-item\" data-id=\"{{ .ID }}\">\n      <a href=\"/{{ .ID }}{{ if $.Data.AppendExt }}{{ .Ext }}{{ end }}\" class=\"upload-link\">{{ if .HasThumb }}<img src=\"/-/thumb/{{ .ID }}.jpg\">{{ else }}<img src=\"/-/static/file.svg\"><div class=\"file-ext-overlay\">{{ .Ext }}</div>{{ end }}</a>\n      <div class=\"history-item-name\" title=\"{{ .Name }}\">{{ .Name }}</div>\n      <div class=\"history-item-data\">{{ .Size }} / <span title=\"{{ .Uploaded.Format \"2006-01-02 15:04:05 MST\" }}\">{{ .Ago }}</span> ago</div>\n      <div class=\"history-item-data\"><a href=\"javascript:\" class=\"delete-upload\">Delete</a></div>\n    </li>\n    {{ end }}\n  </ul>\n  {{ template \"pagination\" . }}\n</section>\n<script>\n  window.addEventListener('DOMContentLoaded', function() {\n    var items = $$('.history-item');\n    for (var i = 0, item; item = items[i]; i++) {\n      (function(item) {\n        item.querySelector('a.delete-upload').addEventListener('click', function() {\n          item.style.opacity = '0.5';\n          var x = new XMLHttpRequest();\n\n          x.addEventListener('load', function(e) {\n            if (e.target.status == 204) {\n              item.style.opacity = '0.0';\n              item.addEventListener('transitionend', function(e) {\n                e.target.parentNode.removeChild(e.target);\n                window.location.reload(true);\n              }, false);\n            } else {\n              item.style.opacity = '';\n              var resp = JSON.parse(x.responseText);\n              if (resp != null && resp.Err != null) {\n                alert('error: ' + resp.Err);\n              } else {\n                alert('wat');\n              }\n            }\n          });\n\n          x.open('POST', '/delete/' + item.dataset.id, true);\n          x.send();\n        }, false);\n      })(item);\n    }\n  }, true);\n</script>\n{{ end }}\n{{ end }}\n\n{{ define \"pagination\" }}\n<nav class=\"pagination\">\n  <span class=\"prevnext{{ if gt .CurrentPage 1 }} active{{ end }}\"><a href=\"/-/history/{{ .PrevPage }}\">Back</a> \xe2\x80\x94</span>\n  Page {{ .CurrentPage }} of {{ .TotalPages }}\n  <span class=\"prevnext{{ if ne .NextPage 0 }} active{{ end }}\">\xe2\x80\x94 <a href=\"/-/history/{{ .NextPage }}\">Next</a></span>\n</nav>\n{{ end }}\n"))
	bindata.RegisterFile("templates/index.tmpl", time.Unix(1442160572, 0), []byte("{{ define \"index\" }}\n  <section id=\"upload\" class=\"floating-section\">\n    <input type=\"file\" id=\"picker\" name=\"picker[]\" multiple>\n    <div id=\"drop-zone\">\n      <div class=\"progress-bar\"></div>\n      <div id=\"drop-zone-text\">Click/tap/drop/paste</div>\n    </div>\n    <div id=\"uploaded-urls\">\n      <ul></ul>\n    </div>\n  </section>\n  <script>window.addEventListener('load', setupUploader, false);</script>\n{{ end }}\n"))
	bindata.RegisterFile("templates/layout.tmpl", time.Unix(1442525678, 0), []byte("{{ define \"head\" }}\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n<link rel=\"shortcut icon\" href=\"/-/static/favicon.png\">\n<link rel=\"apple-touch-icon\" sizes=\"76x76\" href=\"/-/static/airlift_76x76.png\">\n<link rel=\"apple-touch-icon\" sizes=\"120x120\" href=\"/-/static/airlift_120x120.png\">\n<link rel=\"apple-touch-icon\" sizes=\"152x152\" href=\"/-/static/airlift_152x152.png\">\n<link rel=\"apple-touch-icon\" sizes=\"180x180\" href=\"/-/static/airlift_180x180.png\">\n<link rel=\"stylesheet\" href=\"/-/static/style.css\">\n<script src=\"/-/static/script.js\"></script>\n{{ end }}\n{{ define \"common\" }}<!doctype html>\n<html>\n  <head>\n    <title>Airlift</title>\n    {{ template \"head\" }}\n  </head>\n  <body>\n    <div id=\"message-box\"></div>\n    <nav id=\"nav\">\n      <a href=\"/\">Upload</a> /\n      <a href=\"/-/history/0\">History</a> /\n      <a href=\"/-/config\">Configure</a> /\n      <a href=\"/-/logout\">Log out</a>\n    </nav>\n    {{ content }}\n    <div id=\"version\">airliftd {{ .Version }}</div>\n  </body>\n</html>\n{{ end }}\n"))
	bindata.RegisterFile("templates/login.tmpl", time.Unix(1442347253, 0), []byte("{{ define \"login\" }}<!doctype html>\n<html>\n  <head>\n    <title>Log in</title>\n    {{ template \"head\" }}\n  </head>\n  <body>\n    <section id=\"section-login\" class=\"floating-section\">\n      <form method=\"post\" action=\"/-/login\" id=\"login\">\n        {{ if . }}<p id=\"message-box\" class=\"bad\">Incorrect password.</p>{{ end }}\n        <label for=\"password\">Password: </label><input name=\"pass\" id=\"password\" type=\"password\" placeholder=\"password\" autofocus required>\n        <hr>\n        <button type=\"submit\" id=\"submit\">Log in</button>\n      </form>\n    </section>\n  </body>\n</html>\n{{ end }}\n"))
	bindata.RegisterFile("templates/twitterbot.tmpl", time.Unix(1440719768, 0), []byte("{{ define \"twitterbot\" }}\n<!doctype html>\n<html>\n  <head>\n    <title>{{ .ID }}</title>\n    <meta name=\"twitter:card\" content=\"summary_large_image\">\n    <meta name=\"twitter:site\" content=\"{{ .Handle }}\">\n    <meta name=\"twitter:title\" content=\"{{ .ID }}: {{ .Name }}\">\n    <meta name=\"twitter:description\" content=\"{{ .Size }} / uploaded {{ .Uploaded.Format \"2 Jan 2006 15:04\" }}\">\n    <meta name=\"twitter:image\" content=\"http://{{ .Host }}/twitterthumb/{{ .ID }}.jpg\">\n  </head>\n  <body>\n    hi twitterbot\n  </body>\n</html>\n{{ end }}\n"))
}
