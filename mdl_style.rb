# Use all rules..
all

# ...except for the ones with an explicit alternative configuration.
exclude_rule 'MD013' # allow for lines of unlimited length.
exclude_rule 'MD033' # allow for inline HTML.
exclude_rule 'MD036' # allow to use emphasis as headers for repeated headers.
exclude_rule 'MD041' # allow for non-header first lines in files (YAML front-matter, etc).