// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

// Code generated by "mdtogo"; DO NOT EDIT.
package commands

var CatShort = `[Alpha] Print Resource Config from a local directory.`
var CatLong = `
[Alpha]  Print Resource Config from a local directory.

  DIR:
    Path to local directory.
`
var CatExamples = `
    # print Resource config from a directory
    kustomize config cat my-dir/

    # wrap Resource config from a directory in an ResourceList
    kustomize config cat my-dir/ --wrap-kind ResourceList --wrap-version config.kubernetes.io/v1alpha1 --function-config fn.yaml

    # unwrap Resource config from a directory in an ResourceList
    ... | kustomize config cat`

var CompletionShort = `Install shell completion.`
var CompletionLong = `
Install shell completion for kustomize commands and flags -- supports bash, fish and zsh.

    kustomize install-completion

Registers the completion command with known shells (e.g. .bashrc, .bash_profile, etc):

    complete -C /Users/USER/go/bin/kustomize kustomize

Because the completion command is embedded in kustomize directly, there is no need to update
it separately from the kustomize binary.

To uninstall shell completion run:

    COMP_UNINSTALL=1 kustomize install-completion`

var CountShort = `[Alpha] Count Resources Config from a local directory.`
var CountLong = `
[Alpha] Count Resources Config from a local directory.

  DIR:
    Path to local directory.
`
var CountExamples = `
    # print Resource counts from a directory
    kustomize config count my-dir/`

var FmtShort = `[Alpha] Format yaml configuration files.`
var FmtLong = `
[Alpha] Format yaml configuration files.

Fmt will format input by ordering fields and unordered list items in Kubernetes
objects.  Inputs may be directories, files or stdin, and their contents must
include both apiVersion and kind fields.

- Stdin inputs are formatted and written to stdout
- File inputs (args) are formatted and written back to the file
- Directory inputs (args) are walked, each encountered .yaml and .yml file
  acts as an input

For inputs which contain multiple yaml documents separated by \n---\n,
each document will be formatted and written back to the file in the original
order.

Field ordering roughly follows the ordering defined in the source Kubernetes
resource definitions (i.e. go structures), falling back on lexicographical
sorting for unrecognized fields.

Unordered list item ordering is defined for specific Resource types and
field paths.

- .spec.template.spec.containers (by element name)
- .webhooks.rules.operations (by element value)
`
var FmtExamples = `
	# format file1.yaml and file2.yml
	kustomize config fmt file1.yaml file2.yml

	# format all *.yaml and *.yml recursively traversing directories
	kustomize config fmt my-dir/

	# format kubectl output
	kubectl get -o yaml deployments | kustomize config fmt

	# format kustomize output
	kustomize build | kustomize config fmt`

var GrepShort = `[Alpha] Search for matching Resources in a directory or from stdin`
var GrepLong = `
[Alpha] Search for matching Resources in a directory or from stdin.

  QUERY:
    Query to match expressed as 'path.to.field=value'.
    Maps and fields are matched as '.field-name' or '.map-key'
    List elements are matched as '[list-elem-field=field-value]'
    The value to match is expressed as '=value'
    '.' as part of a key or value can be escaped as '\.'

  DIR:
    Path to local directory.
`
var GrepExamples = `
    # find Deployment Resources
    kustomize config grep "kind=Deployment" my-dir/

    # find Resources named nginx
    kustomize config grep "metadata.name=nginx" my-dir/

    # use tree to display matching Resources
    kustomize config grep "metadata.name=nginx" my-dir/ | kustomize config tree

    # look for Resources matching a specific container image
    kustomize config grep "spec.template.spec.containers[name=nginx].image=nginx:1\.7\.9" my-dir/ | kustomize config tree`

var MergeShort = `[Alpha] Merge Resource configuration files`
var MergeLong = `
[Alpha] Merge Resource configuration files

Merge reads Kubernetes Resource yaml configuration files from stdin or sources packages and write
the result to stdout or a destination package.

Resources are merged using the Resource [apiVersion, kind, name, namespace] as the key.  If any of
these are missing, merge will default the missing values to empty.

Resources specified later are high-precedence (the source) and Resources specified
earlier are lower-precedence (the destination).

For information on merge rules, run:

	kustomize config docs merge
`
var MergeExamples = `
    cat resources_and_patches.yaml | kustomize config merge > merged_resources.yaml`

var Merge3Short = `[Alpha] Merge diff of Resource configuration files into a destination (3-way)`
var Merge3Long = `
[Alpha] Merge diff of Resource configuration files into a destination (3-way)

Merge3 performs a 3-way merge by applying the diff between 2 sets of Resources to a 3rd set.

Merge3 may be for rebasing changes to a forked set of configuration -- e.g. compute the difference between the original
set of Resources that was forked and an updated set of those Resources, then apply that difference to the fork.

If a field value differs between the ORIGINAL_DIR and UPDATED_DIR, the value from the UPDATED_DIR is taken and applied
to the Resource in the DEST_DIR.

For information on merge rules, run:

	kustomize config docs-merge3
`
var Merge3Examples = `
    kustomize config merge3 --ancestor a/ --from b/ --to c/`

var RunFnsShort = `[Alpha] Reoncile config functions to Resources.`
var RunFnsLong = `
[Alpha] Reconcile config functions to Resources.

run sequentially invokes all config functions in the directory, providing Resources
in the directory as input to the first function, and writing the output of the last
function back to the directory.

The ordering of functions is determined by the order they are encountered when walking the
directory.  To clearly specify an ordering of functions, multiple functions may be
declared in the same file, separated by '---' (the functions will be invoked in the
order they appear in the file).

#### Arguments:

  DIR:
    Path to local directory.

#### Config Functions:

  Config functions are specified as Kubernetes types containing a metadata.configFn.container.image
  field.  This field tells run how to invoke the container.

  Example config function:

	# in file example/fn.yaml
	apiVersion: fn.example.com/v1beta1
	kind: ExampleFunctionKind
	metadata:
	  configFn:
	    container:
	      # function is invoked as a container running this image
	      image: gcr.io/example/examplefunction:v1.0.1
	  annotations:
	    config.kubernetes.io/local-config: "true" # tools should ignore this
	spec:
	  configField: configValue

  In the preceding example, 'kustomize config run example/' would identify the function by
  the metadata.configFn field.  It would then write all Resources in the directory to
  a container stdin (running the gcr.io/example/examplefunction:v1.0.1 image).  It
  would then write the container stdout back to example/, replacing the directory
  file contents.

  See ` + "`" + `kustomize help config docs-fn` + "`" + ` for more details on writing functions.
`
var RunFnsExamples = `
kustomize config run example/`

var SubShort = `[Alpha] Set values on Resources fields by substituting values.`
var SubLong = `
Set values on Resources fields by substituting predefined markers for new values.

` + "`" + `set` + "`" + ` looks for markers specified on Resource fields and substitute a new user defined
value for the existing value.

` + "`" + `set` + "`" + ` maybe be used to:

- edit configuration programmatically from the cli or scripts
- create reusable bundles of configuration

  DIR

    A directory containing Resource configuration.

  NAME

    Optional.  The name of the substitution to perform or display.

  VALUE

    Optional.  The new value to substitute into the field.


To print the possible substitutions for the Resources in a directory, run ` + "`" + `set` + "`" + ` on
a directory -- e.g. ` + "`" + `kustomize config set DIR/` + "`" + `.

#### Tips

- A description of the value may be specified with ` + "`" + `--description` + "`" + `.
- An owner for the field's value may be defined with ` + "`" + `--owned-by` + "`" + `.
- Prevent overriding previous substitutions with ` + "`" + `--override=false` + "`" + `.
- Revert previous substitutions with ` + "`" + `--revert` + "`" + `.
- Create substitutions on Kustomization.yaml's, patches, etc

When overriding or reverting previous substitutions, the description and owner are left
unmodified unless specified with flags.

To create a substitution for a field see: ` + "`" + `kustomize help config set create` + "`" + `
`
var SubExamples = `
  Resource YAML: Name substitution

    # dir/resources.yaml
    ...
    metadata:
        name: PREFIX-app1 # {"substitutions":[{"name":"prefix","marker":"PREFIX-"}]}
    ...
    ---
    ...
    metadata:
        name: PREFIX-app2 # {"substitutions":[{"name":"prefix","marker":"PREFIX-"}]}
    ...

  Show substitutions: Show the possible substitutions

    $ config set dir
    NAME    DESCRIPTION    VALUE     TYPE    COUNT   SUBSTITUTED   OWNER
    prefix   ''            PREFIX-   string   2       false

  Perform substitution: set a new value, owner and description

    $ config set dir prefix "test-" --description "test environment" --owned-by "dev"
    performed 2 substitutions

  Show substitutions: Show the new values

    $ config set dir
    NAME       DESCRIPTION       VALUE    TYPE    COUNT   SUBSTITUTED   OWNER
    prefix   'test environment'   test-   string   2       true          dev

  New Resource YAML:

    # dir/resources.yaml
    ...
    metadata:
      name: test-app1 # {"substitutions":[{"name":"prefix","marker":"PREFIX-","value":"test-"}],"setBy":"dev","description":"test environment"}
    ...
    ---
    ...
    metadata:
      name: test-app2 # {"substitutions":[{"name":"prefix","marker":"PREFIX-","value":"test-"}],"setBy":"dev","description":"test environment"}
    ...

  Revert substitution:

    config set dir prefix --revert
    performed 2 substitutions

    config set dir
    NAME       DESCRIPTION       VALUE    TYPE    COUNT   SUBSTITUTED   OWNER  
    prefix   'test environment'   PREFIX-  string   2       false          dev    `

var SubsetShort = `[Alpha] Create a new substitution for a Resource field`
var SubsetLong = `
Create a new substitution for a Resource field -- recognized by ` + "`" + `kustomize config set` + "`" + `.

  DIR

    A directory containing Resource configuration.

  NAME

    The name of the substitution to create.

  VALUE

    The current value of the field, or a substring of the field.

#### Tips: Picking Good Marker

Substitutions may be defined by directly editing yaml **or** by running ` + "`" + `kustomize config set create` + "`" + `
to create a new substitution.

Given the YAML:
    
    # resource.yaml
    apiVersion: v1
    kind: Service
    metadata:
      ...
    spec:
      ...
      ports:
        ...
      - name: http
        port: 8080
        ...

Create a new set marker:

    # create a substitution for ports
    $ kustomize config set create dir/ http-port 8080 --type "int" --field "port"

Modified YAML:

    # resource.yaml
    apiVersion: v1
    kind: Service
    metadata:
      ...
    spec:
      ...
      ports:
        ...
      - name: http
        port: 8080 # {"substitutions":[{"name":"port","marker":"[MARKER]"}],"type":"int"}
        ...

Change the value using the ` + "`" + `set` + "`" + ` command:

    # change the http-port value to 8081
    $ kustomize config set dir/ http-port 8081

Resources fields with a field name matching ` + "`" + `--field` + "`" + ` and field value matching ` + "`" + `VALUE` + "`" + ` will
have a line comment added marking this field as settable.

Substitution markers may be:

- valid field values (e.g. ` + "`" + `8080` + "`" + ` for a port)
  - Note: ` + "`" + `008080` + "`" + ` would be preferred because it is more recognizable as a marker
- invalid values that adhere to the schema (e.g. ` + "`" + `0000` + "`" + ` for a port)
- values that do not adhere to the schema (e.g. ` + "`" + `[PORT]` + "`" + ` for port)

Markers **SHOULD be clearly identifiable as a marker and either**:

- **adhere to the field schema** -- e.g. use a valid value


    port: 008080 # {"substitutions":[{"name":"port","marker":"008080"}],"type":"int"}

- **be pre-filled in with a value** -- e.g. set the value when setting the marker


    port: 8080 # {"substitutions":[{"name":"port","marker":"[MARKER]","value":"8080""}],"type":"int"}

**Note:** The important thing is that in both cases the Resource configuration may be directly
applied to a cluster and validated by tools without the tool knowing about the substitution
marker.

The difference between the preceding examples is that:

- the former will be shown as ` + "`" + `SUBSTITUTED=false` + "`" + ` (` + "`" + `config sub dir/` + "`" + ` exits non-0)
- the latter with show up as ` + "`" + `SUBSTITUTED=true` + "`" + ` (` + "`" + `config sub dir/` + "`" + ` exits 0)

When choosing the which to use, consider that checks for unsubstituted values MAY be
configured as pre-commit checks -- if you want to these checks to fail if the value
hasn't been substituted, then don't specify a ` + "`" + `value` + "`" + `.

Markers which are invalid field values MAY be chosen in cases where it is preferred to have
the create or update request fail rather than succeed if the substitution has not yet been
performed.

A substitution may be a substring of the full field:

    $ kustomize config set create dir/ app-image-tag v1.0.01 --type "string" --field "image"

    image: gcr.io/example/app:v1.0.1 # {"substitutions":[{"name":"app-image-tag","marker":"[MARKER]","value":"v1.0.1"}]}


A single field value may have multiple substitutions applied to it:

    name: PREFIX-app-SUFFIX # {"substitutions":[{"name":"prefix","marker":"PREFIX-"},{"name":"suffix","marker":"-SUFFIX"}]}

#### Substitution Format

Substitutions are defined as json encoded FieldMeta comments on fields.

FieldMeta Schema read by ` + "`" + `sub` + "`" + `:

    {
      "title": "FieldMeta",
      "type": "object",
      "properties": {
        "substitutions": {
          "type": "array",
          "description": "Possible substitutions that may be performed against this field.",
          "items": {
            "type": "object",
            "properties": {
              "name": "Name of the substitution.",
              "marker": "Marker for the value to be substituted.",
              "value": "Current substituted value"
            }   
          }
        },
        "type": {
          "type": "string",
          "description": "The value type.  Defaults to string."
          "enum": ["string", "int", "float", "bool"]
        },
        "description": {
          "type": "string",
          "description": "A description of the field's current value.  Optional."
        },
        "setBy": {
          "type": "string",
          "description": "The current owner of the field.  Optional."
        },
      }
    }
`
var SubsetExamples = `
    # set a substitution for port fields matching "8080"
    kustomize config sub create dir/ port 8080 --type "int" --field port \
         --description "default port used by the app"

    # set a substitution for port fields matching "8080", using "0000" as a marker.
    kustomize config sub dir/ port 8080 --marker "0000" --type "int" \
        --field port --description "default port used by the app"

    # substitute a substring of a field rather than the full field -- e.g. only the
    # image tag, not the full image
    kustomize config sub dir/ app-image-tag v1.0.1 --type "string" --substring \
        --field port --description "current stable release"`

var TreeShort = `[Alpha] Display Resource structure from a directory or stdin.`
var TreeLong = `
[Alpha] Display Resource structure from a directory or stdin.

kustomize config tree may be used to print Resources in a directory or cluster, preserving structure

Args:

  DIR:
    Path to local directory directory.

Resource fields may be printed as part of the Resources by specifying the fields as flags.

kustomize config tree has build-in support for printing common fields, such as replicas, container images,
container names, etc.

kustomize config tree supports printing arbitrary fields using the '--field' flag.

By default, kustomize config tree uses Resource graph structure if any relationships between resources (ownerReferences)
are detected, as is typically the case when printing from a cluster. Otherwise, directory graph structure is used. The
graph structure can also be selected explicitly using the '--graph-structure' flag.
`
var TreeExamples = `
    # print Resources using directory structure
    kustomize config tree my-dir/

    # print replicas, container name, and container image and fields for Resources
    kustomize config tree my-dir --replicas --image --name

    # print all common Resource fields
    kustomize config tree my-dir/ --all

    # print the "foo"" annotation
    kustomize config tree my-dir/ --field "metadata.annotations.foo"

    # print the "foo"" annotation
    kubectl get all -o yaml | kustomize config tree \
      --field="status.conditions[type=Completed].status"

    # print live Resources from a cluster using owners for graph structure
    kubectl get all -o yaml | kustomize config tree --replicas --name --image

    # print live Resources with status condition fields
    kubectl get all -o yaml | kustomize config tree \
      --name --image --replicas \
      --field="status.conditions[type=Completed].status" \
      --field="status.conditions[type=Complete].status" \
      --field="status.conditions[type=Ready].status" \
      --field="status.conditions[type=ContainersReady].status"`
