require 'json'

RSpec.describe "planning" do
  let(:plan) { File.read("./tf.json") }
  let(:tf) { json = JSON.parse(plan) }

  def change_null_object()
    {"actions" => [], "before" => nil, "after" => {}, "after_unknown" => {}}
  end

  def get_changed_attributes(path)
    resource = get_changed_resource(path)
    resource ? resource["change"] : change_null_object()
  end

  def get_changed_resource(path)
    return nil if tf["resource_changes"].empty?

    tf["resource_changes"].find {|resource| resource["address"] == path}
  end

  it "registers a temp file path" do
    attrs = get_changed_attributes("module.mymod.local_file.foo")
    expect(attrs["after"]["filename"]).to eq("/tmp/index.md")
  end
end
